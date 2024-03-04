package httpkit

import (
	"context"
	"errors"
	"net/http"
)

type RequestContext struct {
	*RouteHandler

	Writer   http.ResponseWriter
	Response *Response

	Request *http.Request
	Params  Map

	Idx         int
	Middlewares []func(ctx *RequestContext)
}

func (ctx *RequestContext) Next() {
	if ctx.Idx >= len(ctx.Middlewares) {
		panic("end of function chaining")
	}
	ctx.Idx++
	ctx.Middlewares[ctx.Idx](ctx)
}

func (ctx *RequestContext) SendJSON(
	statusCode int, verdict, message string, data interface{}) error {
	if ctx.Response != nil && !ctx.Response.IsEmpty() {
		return errors.New("response already sent")
	}

	ctx.Response = &Response{
		StatusCode: statusCode,
		Verdict:    verdict,
		Message:    message,
		Data:       data,
	}
	return SendJSON(ctx.Writer, statusCode, verdict, message, data)
}

// SendError sends internal error response to client
func (ctx *RequestContext) SendError(err error) error {
	return ctx.SendJSON(http.StatusInternalServerError, VerdictFailure, err.Error(), Map{})
}

// GetContext convenient method to get context from the request
func (ctx *RequestContext) GetContext() context.Context {
	return ctx.Request.Context()
}
