package store

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/trangnkp/my_books/src/model"
	"gorm.io/gorm"
)

type ReadStore struct {
	db *gorm.DB
}

func NewReadStore(db *gorm.DB) *ReadStore {
	return &ReadStore{db: db}
}

func (s *ReadStore) Create(ctx context.Context, record *model.Read) error {
	if err := s.db.WithContext(ctx).Create(record).Error; err != nil {
		log.Errorf("%v", err)
		return err
	}

	return nil
}

type ListReadsFilter struct {
	FromYear int
	ToYear   int
	Source   string
	Language string
}
