package models

type Match struct {
	Id            *int64 `db:"match_id" json:"id,omitempty"`
	PlayerOneId   *int64 `db:"player_one_id" json:"player_one_id,omitempty"`
	PlayerTwoId   *int64 `db:"player_two_id" json:"player_two_id,omitempty"`
	MatchComplete *bool  `db:"match_complete" json:"match_complete,omitempty"`
}
