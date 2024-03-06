package handler

import (
	"github.com/trangnkp/my_books/src/container"
	"github.com/trangnkp/my_books/src/httpkit"
	"net/http"
)

func (app *App) handlePing(ctx *httpkit.RequestContext) {
	_ = ctx.SendJSON(http.StatusOK, httpkit.VerdictSuccess, "pong", container.Map{})
}
