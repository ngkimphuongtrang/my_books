package handler

import (
	"github.com/trangnkp/my_books/src/container"
	"github.com/trangnkp/my_books/src/httpkit"
	"net/http"
	"strconv"
)

type Validation struct {
}

func NewValidation() *Validation {
	return &Validation{}
}

func (v *Validation) validateListParameters(ctx *httpkit.RequestContext) (int64, int64, bool) {
	pageIDParam := ctx.Request.URL.Query().Get("page")
	perPageParam := ctx.Request.URL.Query().Get("per_page")
	var page, perPage int64 = defaultPage, defaultPerPage
	var err error
	if len(pageIDParam) > 0 {
		page, err = strconv.ParseInt(pageIDParam, 10, 64)
		if err != nil {
			_ = ctx.SendJSON(
				http.StatusBadRequest,
				httpkit.VerdictInvalidParameters,
				"offset is not an integer",
				container.Map{})
			return 0, 0, false
		}
	}
	if len(perPageParam) > 0 {
		perPage, err = strconv.ParseInt(perPageParam, 10, 64)
		if err != nil {
			_ = ctx.SendJSON(
				http.StatusBadRequest,
				httpkit.VerdictInvalidParameters,
				"limit is not an integer",
				container.Map{})
			return 0, 0, false
		}
	}
	return page, perPage, true
}
