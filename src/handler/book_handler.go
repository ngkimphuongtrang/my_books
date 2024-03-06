package handler

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/trangnkp/my_books/src/container"
	"github.com/trangnkp/my_books/src/helper"
	"github.com/trangnkp/my_books/src/httpkit"
	"github.com/trangnkp/my_books/src/model"
	"net/http"
	"strings"
)

type CreateBookRequest struct {
	Name   string
	Author string
}

func (app *App) handleCreateBook(ctx *httpkit.RequestContext) {
	var r CreateBookRequest
	if !app.validateParameters(ctx, &r) {
		return
	}
	log.Println(r.Name, r.Author)

	book := &model.Book{Name: r.Name, Author: r.Author}
	err := app.stores.BookStore.Create(ctx.GetContext(), book)
	if err != nil {
		if helper.IsDuplicateKeyError(err) && strings.Contains(err.Error(), "name_author_idx") {
			_ = ctx.SendJSON(
				http.StatusBadRequest,
				httpkit.VerdictExistedRecord,
				"duplicated name or author",
				container.Map{})
			return
		}

		_ = ctx.SendError(err)
		return
	}

	_ = ctx.SendJSON(http.StatusOK, httpkit.VerdictSuccess, "book is created successfully", container.Map{"id": book.ID})
}

func (app *App) validateParameters(ctx *httpkit.RequestContext, r *CreateBookRequest) bool {
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
