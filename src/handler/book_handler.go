package handler

import (
	"encoding/json"
	"github.com/trangnkp/my_books/src/container"
	"github.com/trangnkp/my_books/src/httpkit"
	"github.com/trangnkp/my_books/src/model"
	"net/http"
	"strconv"
)

type CreateBookRequest struct {
	Name   string
	Author string
}

func (app *App) handleCreateBook(ctx *httpkit.RequestContext) {
	var r CreateBookRequest
	if !app.validateCreateBookParameters(ctx, &r) {
		return
	}
	bookID, err := app.stores.BookStore.FindByNameAndAuthor(ctx.GetContext(), r.Name, r.Author)
	if err != nil {
		_ = ctx.SendError(err)
		return
	}
	if bookID > 0 {
		_ = ctx.SendJSON(
			http.StatusBadRequest,
			httpkit.VerdictExistedRecord,
			"duplicated book name and author",
			container.Map{"book_id": bookID})
		return
	}

	book := &model.Book{Name: r.Name, Author: r.Author}
	err = app.stores.BookStore.Create(ctx.GetContext(), book)
	if err != nil {
		_ = ctx.SendError(err)
		return
	}

	_ = ctx.SendJSON(http.StatusOK, httpkit.VerdictSuccess, "book is created successfully", container.Map{"id": book.ID})
}

func (app *App) validateCreateBookParameters(ctx *httpkit.RequestContext, r *CreateBookRequest) bool {
	err := json.NewDecoder(ctx.Request.Body).Decode(r)
	if err != nil {
		_ = ctx.SendJSON(
			http.StatusBadRequest,
			httpkit.VerdictInvalidParameters,
			"parameters are invalid",
			container.Map{})
		return false
	}
	var missingParams []string
	if r.Name == "" {
		missingParams = append(missingParams, "name")
	}
	if r.Author == "" {
		missingParams = append(missingParams, "author")
	}
	if len(missingParams) > 0 {
		_ = ctx.SendJSON(
			http.StatusBadRequest,
			httpkit.VerdictMissingParameters,
			"some required parameters are missing",
			container.Map{"missing_parameters": missingParams})
		return false
	}

	//TODO: standardize parameters
	return true
}

func (app *App) handleListBooks(ctx *httpkit.RequestContext) {
	pageIDParam := ctx.Request.URL.Query().Get("page_id")
	perPageParam := ctx.Request.URL.Query().Get("per_page")
	if len(pageIDParam) == 0 || len(perPageParam) == 0 {
		_ = ctx.SendJSON(
			http.StatusBadRequest,
			httpkit.VerdictInvalidParameters,
			"offset and limit are required to be greater than 0",
			container.Map{})
		return
	}

	pageID, err := strconv.ParseInt(pageIDParam, 10, 64)
	if err != nil {
		_ = ctx.SendJSON(
			http.StatusBadRequest,
			httpkit.VerdictInvalidParameters,
			"offset is not an integer",
			container.Map{})
		return
	}
	perPage, err := strconv.ParseInt(perPageParam, 10, 64)
	if err != nil {
		_ = ctx.SendJSON(
			http.StatusBadRequest,
			httpkit.VerdictInvalidParameters,
			"limit is not an integer",
			container.Map{})
		return
	}
	search := ctx.Request.URL.Query().Get("search")
	offset, limit := (pageID-1)*perPage, perPage
	books, err := app.stores.BookStore.List(ctx.GetContext(), int(offset), int(limit), search)
	if err != nil {
		_ = ctx.SendError(err)
		return
	}

	count, err := app.stores.BookStore.Count(ctx.GetContext(), search)
	if err != nil {
		_ = ctx.SendError(err)
		return
	}

	_ = ctx.SendJSON(http.StatusOK, httpkit.VerdictSuccess, "get list of books successfully",
		container.Map{
			"items": books,
			"count": count,
		})
}
