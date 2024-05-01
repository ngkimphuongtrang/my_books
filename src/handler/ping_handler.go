package handler

import (
	"net/http"

	_const "github.com/ngkimphuongtrang/runkit/const"
	"github.com/ngkimphuongtrang/runkit/container"
	"github.com/ngkimphuongtrang/runkit/httpkit"
)

func (app *App) handlePing(ctx *httpkit.RequestContext) {
	_ = ctx.SendJSON(http.StatusOK, _const.VerdictSuccess, "pong", container.Map{})
}
