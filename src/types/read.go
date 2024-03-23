package types

import (
	"time"

	"github.com/trangnkp/my_books/src/store"
)

type CreateReadRequest struct {
	BookID       int64      `json:"book_id"`
	Source       string     `json:"source"`
	Language     string     `json:"language"`
	FinishedDate *time.Time `json:"finished_date"`
}

func (r *CreateReadRequest) GetMissingParams() []string {
	var missingParams []string
	if r.BookID == 0 {
		missingParams = append(missingParams, "book_id")
	}
	if r.Source == "" {
		missingParams = append(missingParams, "source")
	}
	if r.FinishedDate == nil {
		missingParams = append(missingParams, "finished_date")
	}
	return missingParams
}

func (r *CreateReadRequest) HasValidSource() bool {
	return IsValidSource(r.Source)
}

func IsValidSource(s string) bool {
	for _, source := range store.ReadSources {
		if s == source.String() {
			return true
		}
	}
	return false
}
