package handler

import (
	"net/http"

	"github.com/trangnkp/my_books/src/internal/container"
	"github.com/trangnkp/my_books/src/internal/httpkit"
)

func (app *App) handlePing(ctx *httpkit.RequestContext) {
	_ = ctx.SendJSON(http.StatusOK, httpkit.VerdictSuccess, "pong", container.Map{})
}
