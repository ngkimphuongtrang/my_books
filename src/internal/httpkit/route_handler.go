package httpkit

import "github.com/trangnkp/my_books/src/internal/container"

type RouteHandler struct {
	Route       *Route
	Handle      func(ctx *RequestContext)
	RouteOption container.Map
}
