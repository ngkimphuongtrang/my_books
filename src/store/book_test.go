package store

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/trangnkp/my_books/src/config"
	"github.com/trangnkp/my_books/src/db"
	"github.com/trangnkp/my_books/src/helper"
	"github.com/trangnkp/my_books/src/model"
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
		filter   ListBooksFilter
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
			minCount: 4,
			maxCount: 10,
		},
		{
			name:     "valid_search",
			limit:    5,
			filter:   ListBooksFilter{Name: "nhại"},
			minCount: 1,
			maxCount: 2,
		},
		{
			name:     "2_characters",
			limit:    5,
			filter:   ListBooksFilter{Name: "ch"},
			minCount: 3,
			maxCount: 5,
		},
		{
			name:     "search_by_book_name",
			limit:    10,
			filter:   ListBooksFilter{Name: "giet con chim nhai"},
			minCount: 1,
			maxCount: 5,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			books, err := dbStores.BookStore.List(ctx, tc.offset, tc.limit, &tc.filter)
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

func TestBookStore_Count(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		n := 10
		for i := 0; i < n; i++ {
			book := &model.Book{
				Name: fmt.Sprintf("abc_%d", i),
			}
			err := dbStores.BookStore.Create(ctx, book)
			require.NoError(t, err)
		}

		count, err := dbStores.BookStore.Count(ctx, &ListBooksFilter{Name: "abc"})
		require.NoError(t, err)
		require.LessOrEqual(t, int64(n), count)
	})
}

func TestBookStore_FindByNameAndAuthor(t *testing.T) {
	t.Parallel()

	book := &model.Book{
		Name:   "Atomic Habits",
		Author: "Harper Lee",
	}
	ctx := context.Background()
	err := dbStores.BookStore.Create(ctx, book)
	require.NoError(t, err)

	testCases := []struct {
		name     string
		bookName string
		author   string
		found    bool
	}{
		{
			name:     "found",
			bookName: "Atomic Habits",
			author:   "Harper Lee",
			found:    true,
		},
		{
			name:     "not_found",
			bookName: "Ngay xua co mot con bo",
			found:    false,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			id, err := dbStores.BookStore.FindByNameAndAuthor(ctx, tc.bookName, tc.author)
			require.NoError(t, err)
			if tc.found {
				require.NotZero(t, id)
			} else {
				require.Zero(t, id)
			}
		})
	}
}
