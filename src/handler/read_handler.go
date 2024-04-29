package handler

import (
	"encoding/json"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/trangnkp/my_books/src/internal/container"
	"github.com/trangnkp/my_books/src/internal/httpkit"
	"github.com/trangnkp/my_books/src/model"
	"github.com/trangnkp/my_books/src/serverenv"
	"github.com/trangnkp/my_books/src/service"
	"github.com/trangnkp/my_books/src/store"
	"github.com/trangnkp/my_books/src/types"
)

type ReadHandler struct {
	stores              *store.DBStores
	validation          *Validation
	notificationService *service.KafkaNotificationService
}

func NewReadHandler(env *serverenv.ServerEnv, validation *Validation) *ReadHandler {
	return &ReadHandler{
		stores:              env.DBStores,
		validation:          validation,
		notificationService: env.NotificationService,
	}
}

func (h *ReadHandler) handleCreateRead(ctx *httpkit.RequestContext) {
	var r types.CreateReadRequest
	if !h.validateCreateReadParameters(ctx, &r) {
		return
	}

	if r.Language == "" {
		r.Language = store.LangVI.String()
	}
	bookID := h.getBookID(ctx, &r)
	if bookID == 0 {
		return
	}

	read := &model.Read{BookID: bookID, Source: r.Source, Language: r.Language, FinishedDate: r.FinishedDate}
	err := h.stores.ReadStore.Create(ctx.GetContext(), read)
	if err != nil {
		_ = ctx.SendError(err)
		return
	}

	_ = ctx.SendJSON(http.StatusOK, httpkit.VerdictSuccess, "read is created successfully", container.Map{"id": read.ID})

	err = h.sendEmail("read is created successfully")
	if err != nil {
		log.Error(err)
	}
}

func (h *ReadHandler) getBookID(ctx *httpkit.RequestContext, r *types.CreateReadRequest) int64 {
	book, err := h.stores.BookStore.FindByID(ctx.GetContext(), r.BookID)
	if err != nil {
		_ = ctx.SendError(err)
		return 0
	}
	if book == nil {
		_ = ctx.SendJSON(http.StatusNotFound, httpkit.VerdictRecordNotFound, "book_id not found", container.Map{})
		return 0
	}
	return r.BookID
}

func (h *ReadHandler) validateCreateReadParameters(ctx *httpkit.RequestContext, r *types.CreateReadRequest) bool {
	err := json.NewDecoder(ctx.Request.Body).Decode(r)
	if err != nil {
		_ = ctx.SendJSON(
			http.StatusBadRequest,
			httpkit.VerdictInvalidParameters,
			"parameters are invalid",
			container.Map{"error": err.Error()})
		return false
	}
	missingParams := r.GetMissingParams()
	if len(missingParams) > 0 {
		_ = ctx.SendJSON(
			http.StatusBadRequest,
			httpkit.VerdictMissingParameters,
			"some required parameters are missing",
			container.Map{"missing_parameters": missingParams})
		return false
	}

	if !r.HasValidSource() {
		_ = ctx.SendJSON(
			http.StatusBadRequest,
			httpkit.VerdictInvalidParameters,
			"source is invalid",
			container.Map{"valid_sources": store.ReadSources})
		return false
	}

	//TODO: how language is standardized
	return true
}

func (h *ReadHandler) handleListReads(ctx *httpkit.RequestContext) {
	pageID, perPage, valid := h.validation.validateListParameters(ctx)
	if !valid {
		return
	}
	readFilter, valid := h.getReadFilter(ctx)
	if !valid {
		return
	}

	offset, limit := (pageID-1)*perPage, perPage
	reads, err := h.stores.ReadStore.List(ctx.GetContext(), offset, limit, readFilter)
	if err != nil {
		_ = ctx.SendError(err)
		return
	}

	count, err := h.stores.ReadStore.Count(ctx.GetContext(), readFilter)
	if err != nil {
		_ = ctx.SendError(err)
		return
	}

	_ = ctx.SendJSON(http.StatusOK, httpkit.VerdictSuccess, "list reads successfully",
		container.Map{
			"items": reads,
			"count": count,
		})
}

func (h *ReadHandler) getReadFilter(ctx *httpkit.RequestContext) (*store.ListReadsFilter, bool) {
	fromYearParam := ctx.Request.URL.Query().Get("from_year")
	toYearParam := ctx.Request.URL.Query().Get("to_year")
	language := ctx.Request.URL.Query().Get("language")
	source := ctx.Request.URL.Query().Get("source")
	if len(source) > 0 && !types.IsValidSource(source) {
		_ = ctx.SendJSON(
			http.StatusBadRequest,
			httpkit.VerdictInvalidParameters,
			"source is invalid",
			container.Map{"valid_sources": store.ReadSources})
		return nil, false
	}
	var fromYear, toYear int
	if len(fromYearParam) > 0 {
		fromYearDate, err := time.Parse("2006", fromYearParam)
		if err != nil {
			_ = ctx.SendJSON(
				http.StatusBadRequest,
				httpkit.VerdictInvalidParameters,
				"from_year is in invalid format",
				container.Map{"required_format": "2006"})
			return nil, false
		}
		fromYear = fromYearDate.Year()
	}
	if len(toYearParam) > 0 {
		toYearDate, err := time.Parse("2006", toYearParam)
		if err != nil {
			_ = ctx.SendJSON(
				http.StatusBadRequest,
				httpkit.VerdictInvalidParameters,
				"to_year is in invalid format",
				container.Map{"required_format": "2006"})
			return nil, false
		}
		toYear = toYearDate.Year()
	}

	return &store.ListReadsFilter{
		FromYear: fromYear,
		ToYear:   toYear,
		Language: language,
		Source:   source,
	}, true
}

func (h *ReadHandler) sendEmail(content string) error {
	rawPayload, err := json.Marshal(content)
	if err != nil {
		log.Error(err)
		return err
	}

	if _, err = h.notificationService.Producer.Produce(h.notificationService.Producer.Topic(), []byte("nkpt"), rawPayload); err != nil {
		log.Error(err)
		return err
	}

	log.Info("send email successfully")
	return nil
}
