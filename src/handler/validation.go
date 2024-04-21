package handler

import (
	"net/http"
	"strconv"

	"github.com/trangnkp/my_books/src/internal/container"
	"github.com/trangnkp/my_books/src/internal/httpkit"
)

const (
	defaultPage    = 1
	defaultPerPage = 30
)

type Validation struct {
}

func NewValidation() *Validation {
	return &Validation{}
}

func (v *Validation) validateListParameters(ctx *httpkit.RequestContext) (offset, limit int, ok bool) {
	page, ok := v.parsePaginationParams(ctx, "page", defaultPage)
	if !ok {
		return 0, 0, false // Early return if parsing failed
	}
	perPage, ok := v.parsePaginationParams(ctx, "per_page", defaultPerPage)
	if !ok {
		return 0, 0, false // Early return if parsing failed
	}
	if page < 1 || perPage < 1 {
		_ = ctx.SendJSON(
			http.StatusBadRequest,
			httpkit.VerdictInvalidParameters,
			"page and per_page must be positive integers",
			container.Map{})
		return 0, 0, false
	}
	return page, perPage, true
}

// parsePaginationParams is a helper method to parse pagination parameters.
func (v *Validation) parsePaginationParams(ctx *httpkit.RequestContext, paramName string, defaultValue int) (int, bool) {
	paramValue := ctx.Request.URL.Query().Get(paramName)
	if len(paramValue) > 0 {
		parsedValue, err := strconv.Atoi(paramValue)
		if err != nil {
			_ = ctx.SendJSON(
				http.StatusBadRequest,
				httpkit.VerdictInvalidParameters,
				paramName+" is not an integer",
				container.Map{})
			return 0, false
		}
		return parsedValue, true
	}
	return defaultValue, true
}
