package models

import (
	"database/sql/driver"
	"encoding/json"
	// "fmt"
)

type Turn struct {
	Id         *int64 `db:"turn_id" json:"id,omitempty"`
	MatchId    *int64 `db:"match_id" json:"match_id,omitempty"`
	PlayerId   *int64 `db:"player_id" json:"player_id,omitempty"`
	TurnNumber *int64 `db:"turn_number" json:"turn_number,omitempty"`
	Board      *Json  `db:"gameboard" json:"gameboard,omitempty"`
}

type Json []map[string]interface{}

func (this *Json) Value() (driver.Value, error) {
	jsonBytes, err := json.Marshal(*this)
	return string(jsonBytes), err
}

func (this *Json) Scan(src interface{}) error {
	var jsonMap []map[string]interface{}
	err := json.Unmarshal([]byte(src.([]uint8)), &jsonMap)
	*this = jsonMap

	return err
}
