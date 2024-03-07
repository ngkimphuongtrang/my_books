package store

import (
	"context"
	"errors"
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

func (s *BookStore) FindByID(ctx context.Context, id int64) (*model.Book, error) {
	b := model.Book{}
	err := s.db.WithContext(ctx).Where("id = ?", id).Last(&b).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		log.Errorf("%v", err)
		return nil, err
	}

	return &b, err
}
