package models

type User struct {
	Id       int64      `db:"bb_user_id" json:"id,omitempty"`
	Email    NullString `db:"bb_user_email_address" json:"email,omitempty"`
	Password NullString `db:"bb_user_password" json:"-"`
}
