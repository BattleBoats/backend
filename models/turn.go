package models

type Turn struct {
	Id       *int64     `db:"bb_turn_id" json:"id,omitempty"`
	MatchId  *int64     `db:"bb_turn_match_id" json:"match_id,omitempty"`
	PlayerId *int64     `db:"bb_turn_player_id" json:"player_id,omitempty"`
	Board    *GameBoard `db:"bb_turn_gameboard" json:"gameboard,omitempty"`
}
