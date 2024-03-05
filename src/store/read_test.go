package store

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/trangnkp/my_books/src/config"
	"github.com/trangnkp/my_books/src/db"
	"github.com/trangnkp/my_books/src/model"
	"testing"
	"time"
)

func TestReadStore_Create(t *testing.T) {
	t.Parallel()

	cfg := config.NewDBConfig()
	gormDB, err := db.ConnectORM(cfg)
	require.NoError(t, err)

	t.Run("create", func(t *testing.T) {
		t.Parallel()

		record := &model.Read{
			BookID:       1,
			Source:       SourceAudio.String(),
			Language:     LangVI.String(),
			FinishedDate: time.Now(),
		}

		err := dbStores.ReadStore.Create(context.Background(), record)
		require.NoError(t, err)
		require.NotZero(t, record.ID)
		require.NotZero(t, record.CreatedAt)
		require.NotZero(t, record.UpdatedAt)

		var foundRead *model.Read
		err = gormDB.Take(&foundRead, "id = ?", record.ID).Error
		require.NoError(t, err)

		require.Equal(t, record.BookID, foundRead.BookID)
		require.Equal(t, record.Source, foundRead.Source)
		require.Equal(t, record.Language, foundRead.Language)
	})
}
