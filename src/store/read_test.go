package store

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/trangnkp/my_books/src/config"
	"github.com/trangnkp/my_books/src/db"
	"github.com/trangnkp/my_books/src/model"
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
			FinishedDate: model.NewDate(2024, 3, 24),
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

	record := &model.Read{
		BookID:       1,
		Source:       "hard_copy",
		FinishedDate: model.NewDate(2014, 3, 24),
	}
	err := dbStores.ReadStore.Create(context.Background(), record)
	require.NoError(t, err)

	testCases := []struct {
		name     string
		offset   int
		limit    int
		filter   *ListReadsFilter
		minCount int
		maxCount int
		hasErr   bool
	}{
		{
			name:     "invalid_limit",
			filter:   &ListReadsFilter{},
			minCount: 0,
			maxCount: 0,
			hasErr:   true,
		},
		{
			name:   "hard_copy_source",
			offset: 0,
			limit:  5,
			filter: &ListReadsFilter{
				Source: "hard_copy",
			},
			minCount: 1,
			maxCount: 10,
		},
		{
			name:   "to_year",
			offset: 0,
			limit:  5,
			filter: &ListReadsFilter{
				ToYear: 2014,
			},
			minCount: 1,
			maxCount: 1,
		},
		{
			name:   "from_year",
			offset: 0,
			limit:  5,
			filter: &ListReadsFilter{
				FromYear: 2014,
			},
			minCount: 1,
			maxCount: 3,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			reads, err := dbStores.ReadStore.List(context.Background(), tc.offset, tc.limit, tc.filter)
			if tc.hasErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.LessOrEqual(t, tc.minCount, len(reads))
			require.LessOrEqual(t, len(reads), tc.maxCount)
			for _, read := range reads {
				require.NotNil(t, read.Book)
			}
		})
	}
}

func TestReadStore_Count(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		books, err := dbStores.ReadStore.List(ctx, 0, 100, &ListReadsFilter{})
		require.NoError(t, err)
		n := len(books)
		for i := 0; i < n; i++ {
			read := &model.Read{
				BookID:       books[i].ID,
				FinishedDate: model.NewDate(2024, 3, 24),
			}
			err = dbStores.ReadStore.Create(ctx, read)
			require.NoError(t, err)
		}

		count, err := dbStores.ReadStore.Count(ctx, &ListReadsFilter{})
		require.NoError(t, err)
		require.LessOrEqual(t, int64(n), count)
	})
}
