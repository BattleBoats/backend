package dao

import (
	"fmt"

	"backend/models"
)

const (
	kTURN_TABLE  = "bb_turn"
	kTURN_ID     = "turn_id"
	kTURN_NUMBER = "turn_number"
	kGAMEBOARD   = "gameboard"

	kTURNS_QUERY = "SELECT match_id, turn_id, turn_number, player_id, gameboard FROM %v NATURAL JOIN %v WHERE %v=$1 AND (%v = $2 OR %v = $3) ORDER BY %v ASC"
)

func GetTurns(matchId string, playerId string) ([]*models.Turn, error) {
	dbMap, err := getDbMap()
	defer dbMap.Db.Close()
	if err != nil {
		return nil, err
	}

	var turns []*models.Turn
	query := fmt.Sprintf(kTURNS_QUERY, kTURN_TABLE, kMATCH_TABLE, kMATCH_ID, kMATCH_PLAYER_ONE_ID, kMATCH_PLAYER_TWO_ID, kTURN_NUMBER)
	_, err = dbMap.Select(&turns, query, matchId, playerId, playerId)
	if err != nil {
		return nil, err
	}

	return turns, nil
}

func GetTurn(turnId string, playerId string) (*models.Turn, error) {
	dbMap, err := getDbMap()
	defer dbMap.Db.Close()
	if err != nil {
		return nil, err
	}

	turn := &models.Turn{}
	query := fmt.Sprintf(kTURNS_QUERY, kTURN_TABLE, kMATCH_TABLE, kTURN_ID, kMATCH_PLAYER_ONE_ID, kMATCH_PLAYER_TWO_ID, kTURN_NUMBER)
	err = dbMap.SelectOne(&turn, query, turnId, playerId, playerId)
	if err != nil {
		return nil, err
	}

	return turn, nil
}

func GetMostRecentTurn(matchId string, playerId string) (*models.Turn, error) {
	//get turns
	turns, err := GetTurns(matchId, playerId)
	if err != nil {
		return nil, err
	}

	//return most recent
	var highestTurnNumber int64
	highestTurnNumber = 0
	var highestTurn *models.Turn
	for _, turn := range turns {
		if *turn.TurnNumber >= highestTurnNumber {
			highestTurnNumber = *turn.TurnNumber
			highestTurn = turn
		}
	}

	return highestTurn, nil
}

func InsertTurn(turn *models.Turn) (*models.Turn, error) {
	dbMap, err := getDbMap()
	defer dbMap.Db.Close()
	if err != nil {
		return nil, err
	}

	err = dbMap.Insert(turn)
	if err != nil {
		return nil, err
	}

	return turn, nil
}
