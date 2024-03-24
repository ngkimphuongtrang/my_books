package model

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"strings"
	"time"
)

type Read struct {
	ID           int64     `json:"id"`
	BookID       int64     `json:"book_id"`
	Book         *Book     `json:"book"`
	Source       string    `json:"source"`
	Language     string    `json:"language"`
	FinishedDate *Date     `json:"finished_date"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Date struct {
	year  int
	month time.Month
	day   int
}

// Value implements the driver.Valuer interface, marshalling to a database value.
func (d *Date) Value() (driver.Value, error) {
	if d == nil {
		return nil, nil
	}

	tt := time.Date(d.year, d.month, d.day, 0, 0, 0, 0, time.UTC)
	return tt.Format(time.DateTime), nil
}

func (d *Date) Scan(src interface{}) error {
	tt, ok := src.(time.Time)
	if !ok {
		return fmt.Errorf("invalid type: %v", reflect.TypeOf(src))
	}

	*d = Date{
		year:  tt.Year(),
		month: tt.Month(),
		day:   tt.Day(),
	}
	return nil
}

// UnmarshalJSON implements the csvutil.Unmarshaler interface, unmarshalling from a CSV value.
func (d *Date) UnmarshalJSON(data []byte) error {
	// Trim the leading and trailing " from the JSON string value.
	strInput := strings.Trim(string(data), `"`)

	// Handle empty date string (which you can interpret as null or zero date)
	if strInput == "null" || strInput == "" {
		*d = Date{} // Set to zero date if null or empty.
		return nil
	}

	// Parse the date string in "YYYY-MM-DD" format
	tt, err := time.Parse("2006-01-02", strInput)
	if err != nil {
		return err
	}

	// Set the parsed values to the Date struct
	*d = Date{
		year:  tt.Year(),
		month: tt.Month(),
		day:   tt.Day(),
	}
	return nil
}

func NewDate(year int, month time.Month, day int) *Date {
	return &Date{year: year, month: month, day: day}
}
