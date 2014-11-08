package models

import (
	"time"
)

type PlayerSession struct {
	Id       string    `json:"id,omitempty"`
	PlayerId string    `json:"player_id,omitempty"`
	Expire   time.Time `json:"expire,omitempty"`
}
