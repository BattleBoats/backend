package dao

import (
	"fmt"

	"backend/models"
)

const (
	kMATCH_TABLE         = "bb_match"
	kMATCH_ID            = "match_id"
	kMATCH_PLAYER_ONE_ID = "player_one_id"
	kMATCH_PLAYER_TWO_ID = "player_two_id"
	kMATCH_COMPLETE      = "match_complete"
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

func GetMatches(playerId string, complete bool) ([]*models.Match, error) {
	dbMap, err := getDbMap()
	defer dbMap.Db.Close()
	if err != nil {
		return nil, err
	}

	var matches []*models.Match
	query := fmt.Sprintf("SELECT * FROM %v WHERE (%v=$1 OR %v=$2) AND %v=$3", kMATCH_TABLE, kMATCH_PLAYER_ONE_ID, kMATCH_PLAYER_TWO_ID, kMATCH_COMPLETE)
	_, err = dbMap.Select(&matches, query, playerId, playerId, complete)
	if err != nil {
		return nil, err
	}

	return matches, nil
}

// func GetIncompleteMatches(playerId string) ([]*models.Match, error) {
// 	return nil, nil
// }

func InsertMatch(match *models.Match) (*models.Match, error) {
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
