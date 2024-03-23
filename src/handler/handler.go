package handler

import (
	"github.com/trangnkp/my_books/src/store"
)

type Handler struct {
	*BookHandler
	*ReadHandler
}

func NewHandler(stores *store.DBStores) *Handler {
	return &Handler{
		BookHandler: NewBookHandler(stores),
		ReadHandler: NewReadHandler(stores),
	}
}
