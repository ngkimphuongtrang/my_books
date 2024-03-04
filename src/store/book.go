package store

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/trangnkp/my_books/src/model"
	"gorm.io/gorm"
)

type BookStore struct {
	db *gorm.DB
}

func NewBookStore(db *gorm.DB) *BookStore {
	return &BookStore{db: db}
}

func (s *BookStore) Create(ctx context.Context, record *model.Book) error {
	if err := s.db.WithContext(ctx).Create(record).Error; err != nil {
		log.Errorf("%v", err)
		return err
	}
	return nil
}
