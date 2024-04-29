package handler

import (
	"context"
	"os"
	"testing"

	"github.com/trangnkp/my_books/src/config"
	"github.com/trangnkp/my_books/src/model"
	"github.com/trangnkp/my_books/src/serverenv"
	"github.com/trangnkp/my_books/src/store"
	"github.com/trangnkp/my_books/src/util"
)

var testApp *App
var bookHandler *BookHandler
var readHandler *ReadHandler

func TestMain(m *testing.M) {
	os.Exit(testMainWrapper(m))
}

func testMainWrapper(m *testing.M) int {
	cfg := config.New()

	stores, err := store.NewDBStores(cfg.DB)
	if err != nil {
		panic(err)
	}

	err = stores.Reset(util.GetProjectRoot())
	if err != nil {
		panic(err)
	}

	env, err := serverenv.NewServerEnv(cfg)
	if err != nil {
		panic(err)
	}
	seedData(stores)
	bookHandler = NewBookHandler(stores, nil)
	readHandler = NewReadHandler(env, nil)

	testApp = NewApp(cfg)
	return m.Run()
}

func seedData(stores *store.DBStores) {
	record := &model.Book{
		Name:   "duplicated",
		Author: "duplicated",
	}
	err := stores.BookStore.Create(context.Background(), record)
	if err != nil {
		panic(err)
	}
}
