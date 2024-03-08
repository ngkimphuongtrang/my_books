package store

import (
	"context"
	"errors"
	"fmt"
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

func (s *BookStore) FindByNameAndAuthor(ctx context.Context, name, author string) (int64, error) {
	b := model.Book{}
	err := s.db.WithContext(ctx).
		Where("name = ? AND author = ?", name, author).Last(&b).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, nil
		}
		log.Errorf("%v", err)
		return 0, err
	}
	return b.ID, nil
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

func (s *BookStore) List(ctx context.Context, offset, limit int, search string) ([]*model.Book, error) {
	if offset < 0 || limit <= 0 {
		err := fmt.Errorf("invalid offset or limit, offset=%d, limit=%d", offset, limit)
		log.Errorf("%v", err)
		return nil, err
	}

	db := s.db.WithContext(ctx)
	if len(search) > 0 {
		db = db.Where("MATCH (name) AGAINST (? IN BOOLEAN MODE)", search)
	}

	var books []*model.Book
	err := db.Offset(offset).Limit(limit).Order("id ASC").Find(&books).Error
	if err != nil {
		log.Errorf("%v", err)
		return nil, err
	}

	return books, nil
}

func (s *BookStore) Count(ctx context.Context, search string) (int64, error) {
	db := s.db.WithContext(ctx)
	if len(search) > 0 {
		db = db.Where("MATCH (name) AGAINST (? IN BOOLEAN MODE)", search)
	}

	var count int64
	err := db.Model(&model.Book{}).Count(&count).Error
	if err != nil {
		log.Errorf("%v", err)
		return 0, err
	}

	return count, nil
}
