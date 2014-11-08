package dao

import (
	"fmt"

	"backend/models"
)

const (
	kMATCH_TABLE         = "bb_match"
	kMATCH_ID            = "bb_match_id"
	kMATCH_PLAYER_ONE_ID = "bb_match_player_one_id"
	kMATCH_PLAYER_TWO_ID = "bb_match_player_two_id"
)

func GetMatchById(matchId string) (*models.Match, error) {
	dbMap, err := getDbMap()
	defer dbMap.Db.Close()
	if err != nil {
		return nil, err
	}

	match := &models.Match{}
	query := fmt.Sprintf("SELECT * FROM %v WHERE %v=$1", kMATCH_TABLE, kMATCH_ID)
	err = dbMap.SelectOne(&match, query, matchId)
	if err != nil {
		return nil, err
	}

	return match, nil
}

// func GetAvailableMatch() (*models.Match, error) {
// 	dbMap, err := getDbMap()
// 	defer dbMap.Db.Close()
// 	if err != nil {
// 		return nil, err
// 	}

// 	match := &models.Match{}
// 	subQuery := fmt.Sprintf("SELECT %v FROM %v WHERE %v IS NULL", kMATCH_ID, kMATCH_TABLE, kMATCH_PLAYER_TWO_ID)
// err = dbMap.SelectOne(&match, query)

// }

func InsertMatch(playerOneId int64, playerTwoId int64) (*models.Match, error) {
	match := &models.Match{
		PlayerOneId: &playerOneId,
		PlayerTwoId: &playerTwoId,
	}

	dbMap, err := getDbMap()
	defer dbMap.Db.Close()
	if err != nil {
		return nil, err
	}

	err = dbMap.Insert(match)
	if err != nil {
		return nil, err
	}

	return match, nil
}
