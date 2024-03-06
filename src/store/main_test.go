package store

import (
	"github.com/trangnkp/my_books/src/config"
	"os"
	"testing"
)

var dbStores *DBStores

func TestMain(m *testing.M) {
	var err error
	dbStores, err = NewDBStores(config.NewDBConfig())
	if err != nil {
		panic(err)
	}

	if err = dbStores.Migrate(""); err != nil {
		panic(err)
	}

	code := m.Run()
	_ = dbStores.Close()
	os.Exit(code)
}
