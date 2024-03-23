package model

import "time"

type Read struct {
	ID           int64     `json:"id"`
	BookID       int64     `json:"book_id"`
	Book         *Book     `json:"book"`
	Source       string    `json:"source"`
	Language     string    `json:"language"`
	FinishedDate time.Time `json:"finished_date"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
