package models

import (
	"time"
)

type UserSession struct {
	Id     string    `json:"id,omitempty"`
	UserId string    `json:"userid,omitempty"`
	Expire time.Time `json:"expire,omitempty"`
}
