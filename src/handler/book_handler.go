package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ngkimphuongtrang/runkit/const"
	"github.com/ngkimphuongtrang/runkit/container"
	"github.com/ngkimphuongtrang/runkit/httpkit"
	"github.com/trangnkp/my_books/src/model"
	"github.com/trangnkp/my_books/src/store"
	"github.com/trangnkp/my_books/src/types"
)

type BookHandler struct {
	stores     *store.DBStores
	validation *Validation
}

func NewBookHandler(stores *store.DBStores, validation *Validation) *BookHandler {
	return &BookHandler{
		stores:     stores,
		validation: validation,
	}
}

func (h *BookHandler) handleCreateBook(ctx *httpkit.RequestContext) {
	var r types.CreateBookRequest
	if !h.validateCreateBookParameters(ctx, &r) {
		return
	}
	bookID, err := h.stores.BookStore.FindByNameAndAuthor(ctx.GetContext(), r.Name, r.Author)
	if err != nil {
		_ = ctx.SendError(err)
		return
	}
	if bookID > 0 {
		_ = ctx.SendJSON(
			http.StatusBadRequest,
			_const.VerdictExistedRecord,
			"duplicated book name and author",
			container.Map{"book_id": bookID})
		return
	}

	book := &model.Book{Name: r.Name, Author: r.Author}
	err = h.stores.BookStore.Create(ctx.GetContext(), book)
	if err != nil {
		_ = ctx.SendError(err)
		return
	}

	_ = ctx.SendJSON(http.StatusOK, _const.VerdictSuccess, "book is created successfully", container.Map{"id": book.ID})
}

func (h *BookHandler) validateCreateBookParameters(ctx *httpkit.RequestContext, r *types.CreateBookRequest) bool {
	err := json.NewDecoder(ctx.Request.Body).Decode(r)
	if err != nil {
		_ = ctx.SendJSON(
			http.StatusBadRequest,
			_const.VerdictInvalidParameters,
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
			_const.VerdictMissingParameters,
			"some required parameters are missing",
			container.Map{"missing_parameters": missingParams})
		return false
	}

	//TODO: standardize parameters
	return true
}

func (h *BookHandler) handleListBooks(ctx *httpkit.RequestContext) {
	pageID, perPage, valid := h.validation.validateListParameters(ctx)
	if !valid {
		return
	}
	search := ctx.Request.URL.Query().Get("search")
	offset, limit := (pageID-1)*perPage, perPage
	filter := &store.ListBooksFilter{Name: search}
	books, err := h.stores.BookStore.List(ctx.GetContext(), offset, limit, filter)
	if err != nil {
		_ = ctx.SendError(err)
		return
	}

	count, err := h.stores.BookStore.Count(ctx.GetContext(), filter)
	if err != nil {
		_ = ctx.SendError(err)
		return
	}

	_ = ctx.SendJSON(http.StatusOK, _const.VerdictSuccess, "get list of books successfully",
		container.Map{
			"items": books,
			"count": count,
		})
}
