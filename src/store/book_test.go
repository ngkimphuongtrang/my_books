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
			Name:   "Who moved my cheese?",
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
