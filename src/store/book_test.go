package store

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/trangnkp/my_books/src/config"
	"github.com/trangnkp/my_books/src/db"
	"github.com/trangnkp/my_books/src/helper"
	"github.com/trangnkp/my_books/src/model"
	"testing"
)

func TestBookStore_Create(t *testing.T) {
	t.Parallel()

	cfg := config.NewDBConfig()
	gormDB, err := db.ConnectORM(cfg)
	require.NoError(t, err)

	t.Run("create", func(t *testing.T) {
		t.Parallel()

		record := &model.Book{
			Name:   "Who moved my xeese?",
			Author: "X",
		}

		err := dbStores.BookStore.Create(context.Background(), record)
		require.NoError(t, err)
		require.NotZero(t, record.ID)
		require.NotZero(t, record.CreatedAt)
		require.NotZero(t, record.UpdatedAt)

		var foundBook *model.Book
		err = gormDB.Take(&foundBook, "id = ?", record.ID).Error
		require.NoError(t, err)

		require.Equal(t, record.Name, foundBook.Name)
		require.Equal(t, record.Author, foundBook.Author)
	})

	t.Run("duplicated book", func(t *testing.T) {
		t.Parallel()

		record := &model.Book{
			Name:   "duplicated_name",
			Author: "X",
		}
		err := dbStores.BookStore.Create(context.Background(), record)
		require.NoError(t, err)

		err = dbStores.BookStore.Create(context.Background(), record)
		require.Error(t, err)
		require.True(t, helper.IsDuplicateKeyError(err))
	})
}

func TestBookStore_List(t *testing.T) {
	t.Parallel()

	books := []*model.Book{
		{
			Name: "Giết con chim nhại",
		},
		{
			Name: "Chiến binh cầu vồng",
		},
		{
			Name: "Who moved my cheese?",
		},
		{
			Name: "Đi tìm lẽ sống",
		},
	}
	ctx := context.Background()
	for _, book := range books {
		err := dbStores.BookStore.Create(ctx, book)
		require.NoError(t, err)
	}

	testCases := []struct {
		name     string
		offset   int
		limit    int
		search   string
		minCount int
		maxCount int
		wantErr  bool
	}{
		{
			name:    "invalid_limit",
			limit:   0,
			wantErr: true,
		},
		{
			name:     "empty_search",
			limit:    5,
			search:   "",
			minCount: 4,
			maxCount: 4,
		},
		{
			name:     "valid_search",
			limit:    5,
			search:   "nhại",
			minCount: 1,
			maxCount: 1,
		},
		{
			name:     "2_characters",
			limit:    5,
			search:   "ch",
			minCount: 3,
			maxCount: 3,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			books, err := dbStores.BookStore.List(ctx, tc.offset, tc.limit, tc.search)
			if tc.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.LessOrEqual(t, tc.minCount, len(books))
			require.LessOrEqual(t, len(books), tc.maxCount)
			for i, book := range books {
				if i > 0 {
					require.Less(t, books[i-1].ID, book.ID)
				}
			}
		})
	}
}
