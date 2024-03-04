package main

import (
	"github.com/trangnkp/my_books/src/config"
	"github.com/trangnkp/my_books/src/handler"
	"github.com/trangnkp/my_books/src/store"
)

func main() {
	cfg := config.New()

	stores, err := store.NewDBStores(cfg.DB)
	if err != nil {
		panic(err)
	}
	app := handler.NewApp(cfg, stores)
	app.Start()
}
