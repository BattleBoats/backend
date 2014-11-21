package models

type Player struct {
	Id       *int64  `db:"player_id" json:"id,omitempty"`
	Email    *string `db:"email_address" json:"email,omitempty"`
	Password *string `db:"password" json:"-"`
}
