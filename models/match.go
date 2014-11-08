package models

type Match struct {
	Id          *int64 `db:"bb_match_id" json:"id,omitempty"`
	PlayerOneId *int64 `db:"bb_match_player_one_id" json:"player_one_id,omitempty"`
	PlayerTwoId *int64 `db:"bb_match_player_two_id" json:"player_two_id,omitempty"`
}
