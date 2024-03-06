package httpkit

import "github.com/trangnkp/my_books/src/container"

type RouteHandler struct {
	Route       *Route
	Handle      func(ctx *RequestContext)
	RouteOption container.Map
}
