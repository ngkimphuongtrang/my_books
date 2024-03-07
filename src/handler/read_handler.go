package handler

import (
	"encoding/json"
	"github.com/trangnkp/my_books/src/container"
	"github.com/trangnkp/my_books/src/httpkit"
	"github.com/trangnkp/my_books/src/model"
	"github.com/trangnkp/my_books/src/store"
	"net/http"
	"strconv"
	"time"
)

type CreateReadRequest struct {
	BookID       string     `json:"book_id"`
	Source       string     `json:"source"`
	Language     string     `json:"language"`
	FinishedDate *time.Time `json:"finished_date"`
}

func (app *App) handleCreateRead(ctx *httpkit.RequestContext) {
	var r CreateReadRequest
	if !app.validateCreateReadParameters(ctx, &r) {
		return
	}

	if r.Language == "" {
		r.Language = store.LangVI.String()
	}
	bookID, err := strconv.ParseInt(r.BookID, 10, 64)
	if err != nil {
		_ = ctx.SendJSON(
			http.StatusBadRequest,
			httpkit.VerdictInvalidParameters,
			"book_id is not an integer",
			container.Map{})
		return
	}
	book, err := app.stores.BookStore.FindByID(ctx.GetContext(), bookID)
	if err != nil {
		_ = ctx.SendError(err)
		return
	}
	if book == nil {
		_ = ctx.SendJSON(http.StatusNotFound, httpkit.VerdictRecordNotFound, "book_id not found", container.Map{})
		return
	}

	read := &model.Read{BookID: bookID, Source: r.Source, Language: r.Language, FinishedDate: *r.FinishedDate}
	err = app.stores.ReadStore.Create(ctx.GetContext(), read)
	if err != nil {
		_ = ctx.SendError(err)
		return
	}

	_ = ctx.SendJSON(http.StatusOK, httpkit.VerdictSuccess, "read is created successfully", container.Map{"id": read.ID})
}

func (app *App) validateCreateReadParameters(ctx *httpkit.RequestContext, r *CreateReadRequest) bool {
	err := json.NewDecoder(ctx.Request.Body).Decode(r)
	if err != nil {
		_ = ctx.SendJSON(
			http.StatusBadRequest,
			httpkit.VerdictInvalidParameters,
			"parameters are invalid",
			container.Map{"error": err.Error()})
		return false
	}
	missingParams := r.getMissingParams()
	if len(missingParams) > 0 {
		_ = ctx.SendJSON(
			http.StatusBadRequest,
			httpkit.VerdictMissingParameters,
			"some required parameters are missing",
			container.Map{"missing_parameters": missingParams})
		return false
	}

	if !r.isValidSource() {
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

func (r *CreateReadRequest) getMissingParams() []string {
	var missingParams []string
	if r.BookID == "" {
		missingParams = append(missingParams, "book_id")
	}
	if r.Source == "" {
		missingParams = append(missingParams, "source")
	}
	if r.FinishedDate == nil {
		missingParams = append(missingParams, "finished_date")
	}
	return missingParams
}

func (r *CreateReadRequest) isValidSource() bool {
	for _, source := range store.ReadSources {
		if r.Source == source.String() {
			return true
		}
	}
	return false
}
