package handler

import (
	"net/http"

	"github.com/ngkimphuongtrang/runkit/httpkit"
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
		Path:   "/books",
	}

	RouteCreateRead = &httpkit.Route{
		Name:   "create_read",
		Method: http.MethodPost,
		Path:   "/reads",
	}

	RouteListBooks = &httpkit.Route{
		Name:   "list_books",
		Method: http.MethodGet,
		Path:   "/books",
	}

	RouteListReads = &httpkit.Route{
		Name:   "list_reads",
		Method: http.MethodGet,
		Path:   "/reads",
	}
)
