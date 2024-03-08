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

func TestCreate_List(t *testing.T) {
	t.Parallel()

	read := &model.Read{
		BookID:       1,
		Source:       "hard_copy",
		FinishedDate: time.Now().AddDate(-10, 0, 0),
	}
	err := dbStores.ReadStore.Create(context.Background(), read)
	require.NoError(t, err)

	testCases := []struct {
		name     string
		filter   *ListReadsFilter
		minCount int
		maxCount int
	}{
		{
			name: "hard_copy_source",
			filter: &ListReadsFilter{
				Source: "hard_copy",
			},
			minCount: 1,
			maxCount: 10,
		},
		{
			name: "from_year",
			filter: &ListReadsFilter{
				ToYear: 2014,
			},
			minCount: 1,
			maxCount: 1,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			reads, err := dbStores.ReadStore.List(context.Background(), 0, 5, tc.filter)
			require.NoError(t, err)
			require.LessOrEqual(t, tc.minCount, len(reads))
			require.LessOrEqual(t, len(reads), tc.maxCount)
		})
	}
}
