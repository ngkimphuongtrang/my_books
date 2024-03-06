package store

import (
	"github.com/trangnkp/my_books/src/config"
	"github.com/trangnkp/my_books/src/util"
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

	err = dbStores.Reset(util.GetProjectRoot())
	if err != nil {
		panic(err)
	}
	
	if err = dbStores.Migrate(util.GetProjectRoot()); err != nil {
		panic(err)
	}

	code := m.Run()
	_ = dbStores.Close()
	os.Exit(code)
}
