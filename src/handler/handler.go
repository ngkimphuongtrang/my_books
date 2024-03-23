package handler

import (
	"github.com/trangnkp/my_books/src/store"
)

type Handler struct {
	*BookHandler
	*ReadHandler
}

func NewHandler(stores *store.DBStores, validation *Validation) *Handler {
	return &Handler{
		BookHandler: NewBookHandler(stores, validation),
		ReadHandler: NewReadHandler(stores, validation),
	}
}
