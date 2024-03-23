package store

import (
	"context"
	"fmt"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
	"github.com/trangnkp/my_books/src/model"
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

func (s *ReadStore) List(ctx context.Context, offset, limit int, filter *ListReadsFilter) ([]*model.Read, error) {
	if offset < 0 || limit <= 0 {
		err := fmt.Errorf("invalid offset or limit, offset=%d, limit=%d", offset, limit)
		log.Errorf("%v", err)
		return nil, err
	}

	db := s.db.WithContext(ctx)
	db = filter.buildQuery(db)

	var reads []*model.Read

	// Preload("Book") is equivalent to join books on reads.book_id = books.id
	err := db.Preload("Book").Offset(offset).Limit(limit).Order("reads.id ASC").Find(&reads).Error
	if err != nil {
		log.Errorf("%v", err)
		return nil, err
	}

	return reads, nil
}

func (s *ReadStore) Count(ctx context.Context, filter *ListReadsFilter) (int64, error) {
	db := s.db.WithContext(ctx)
	db = filter.buildQuery(db)

	var count int64
	err := db.Model(&model.Read{}).Count(&count).Error
	if err != nil {
		log.Errorf("%v", err)
		return 0, err
	}

	return count, nil
}

func (f *ListReadsFilter) buildQuery(db *gorm.DB) *gorm.DB {
	if f.FromYear > 0 {
		db = db.Where("year(finished_date) >= ?", f.FromYear)
	}
	if f.ToYear > 0 {
		db = db.Where("year(finished_date) <= ?", f.ToYear)
	}
	if len(f.Language) > 0 {
		db = db.Where("language = ?", f.Language)
	}
	if len(f.Source) > 0 {
		db = db.Where("source = ?", f.Source)
	}
	return db
}
