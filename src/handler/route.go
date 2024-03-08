package handler

import (
	"github.com/trangnkp/my_books/src/httpkit"
	"net/http"
)

var (
	RoutePing = &httpkit.Route{
		Name:   "ping",
		Method: http.MethodGet,
		Path:   "/ping",
	}

	RouteCreateBook = &httpkit.Route{
		Name:   "create_book",
		Method: http.MethodPost,
		Path:   "/book",
	}

	RouteCreateRead = &httpkit.Route{
		Name:   "create_read",
		Method: http.MethodPost,
		Path:   "/read",
	}

	RouteListBooks = &httpkit.Route{
		Name:   "list_books",
		Method: http.MethodGet,
		Path:   "/books",
	}
)
