package models

type Turn struct {
	Id         *int64    `db:"turn_id" json:"id,omitempty"`
	MatchId    *int64    `db:"match_id" json:"match_id,omitempty"`
	PlayerId   *int64    `db:"player_id" json:"player_id,omitempty"`
	TurnNumber *int64    `db:"turn_number" json:"turn_number,omitempty"`
	Board      GameBoard `db:"gameboard" json:"gameboard,omitempty"`
}
