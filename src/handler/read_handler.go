package handler

import (
	"encoding/json"
	"github.com/trangnkp/my_books/src/container"
	"github.com/trangnkp/my_books/src/httpkit"
	"github.com/trangnkp/my_books/src/model"
	"github.com/trangnkp/my_books/src/store"
	"github.com/trangnkp/my_books/src/types"
	"net/http"
	"time"
)

type ReadHandler struct {
	stores     *store.DBStores
	validation *Validation
}

func NewReadHandler(stores *store.DBStores, validation *Validation) *ReadHandler {
	return &ReadHandler{
		stores:     stores,
		validation: validation,
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

	read := &model.Read{BookID: bookID, Source: r.Source, Language: r.Language, FinishedDate: *r.FinishedDate}
	err := h.stores.ReadStore.Create(ctx.GetContext(), read)
	if err != nil {
		_ = ctx.SendError(err)
		return
	}

	_ = ctx.SendJSON(http.StatusOK, httpkit.VerdictSuccess, "read is created successfully", container.Map{"id": read.ID})
}

func (h *ReadHandler) getBookID(ctx *httpkit.RequestContext, r *types.CreateReadRequest) int64 {
	if r.BookID > 0 {
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
	filter := &store.ListBooksFilter{Name: r.BookName}
	books, err := h.stores.BookStore.List(ctx.GetContext(), 0, 2, filter)
	if err != nil {
		_ = ctx.SendError(err)
		return 0
	}
	if len(books) > 1 {
		_ = ctx.SendJSON(http.StatusBadRequest, httpkit.VerdictUnspecifiedResource, "multiple books found", container.Map{})
		return 0
	}
	if len(books) == 0 {
		_ = ctx.SendJSON(http.StatusNotFound, httpkit.VerdictRecordNotFound, "book_name not found", container.Map{})
		return 0
	}
	return books[0].ID
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

	if r.BookID > 0 && r.BookName != "" {
		_ = ctx.SendJSON(
			http.StatusBadRequest,
			httpkit.VerdictRedundant,
			"book_id and book_name are mutually exclusive",
			container.Map{})
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
	reads, err := h.stores.ReadStore.List(ctx.GetContext(), int(offset), int(limit), readFilter)
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
		toYearDate, err := time.Parse(time.RFC3339, toYearParam)
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
