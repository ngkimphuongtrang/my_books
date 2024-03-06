package handler

import (
	"github.com/trangnkp/my_books/src/config"
	"github.com/trangnkp/my_books/src/store"
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
	if err := stores.Migrate(); err != nil {
		panic(err)
	}

	testApp = NewApp(cfg, stores)
	return m.Run()
}
