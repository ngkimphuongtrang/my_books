package store

import (
	"github.com/trangnkp/my_books/src/db"
	"gorm.io/gorm"
)

type SQLStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *SQLStore {
	return &SQLStore{db: db}
}

type DBStores struct {
	*SQLStore
	config *db.MySQLConfig

	BookStore *BookStore
}

func NewDBStores(cfg *db.MySQLConfig) (*DBStores, error) {
	gormDB, err := db.ConnectORM(cfg)
	if err != nil {
		return nil, err
	}

	return &DBStores{
		config:    cfg,
		SQLStore:  NewSQLStore(gormDB),
		BookStore: NewBookStore(gormDB),
	}, nil
}
