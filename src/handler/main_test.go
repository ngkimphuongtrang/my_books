package handler

import (
	"context"
	"github.com/trangnkp/my_books/src/config"
	"github.com/trangnkp/my_books/src/model"
	"github.com/trangnkp/my_books/src/store"
	"github.com/trangnkp/my_books/src/util"
	"os"
	"testing"
)

var testApp *App

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
	//if err = stores.Migrate(util.GetProjectRoot()); err != nil {
	//	panic(err)
	//}
	seedData(stores)

	testApp = NewApp(cfg, stores)
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
