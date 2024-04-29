package handler

import (
	"github.com/trangnkp/my_books/src/serverenv"
)

type Handler struct {
	*BookHandler
	*ReadHandler
}

func NewHandler(env *serverenv.ServerEnv, validation *Validation) *Handler {
	return &Handler{
		BookHandler: NewBookHandler(env.DBStores, validation),
		ReadHandler: NewReadHandler(env, validation),
	}
}
