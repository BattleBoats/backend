package models

type Player struct {
	Id       *int64  `db:"bb_player_id" json:"id,omitempty"`
	Email    *string `db:"bb_player_email_address" json:"email,omitempty"`
	Password *string `db:"bb_player_password" json:"-"`
}
