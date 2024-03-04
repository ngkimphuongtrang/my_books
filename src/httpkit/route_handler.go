package httpkit

type RouteHandler struct {
	Route       *Route
	Handle      func(ctx *RequestContext)
	RouteOption Map
}
