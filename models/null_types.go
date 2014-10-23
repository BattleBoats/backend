package models

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/lib/pq"
)

type NullInt64 struct {
	sql.NullInt64
}

func (i *NullInt64) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Int64)
}

func BoxInt64(i int64) NullInt64 {
	nullInt := NullInt64{}
	nullInt.Int64 = i
	nullInt.Valid = true
	return nullInt
}

type NullString struct {
	sql.NullString
}

func (s *NullString) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String)
}

func BoxString(s string) NullString {
	nullString := NullString{}
	nullString.String = s
	nullString.Valid = true
	return nullString
}

type NullTime struct {
	pq.NullTime
}

func (t *NullTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Time)
}

func BoxTime(t time.Time) NullTime {
	nullTime := NullTime{}
	nullTime.Time = t
	nullTime.Valid = true
	return nullTime
}
