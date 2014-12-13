package services

import (
	// "fmt"
	// "log"
	// "time"
	// "database/sql"
	"encoding/json"
	"strconv"

	"backend/dao"
	"backend/errors"
	"backend/models"
	// "backend/models/sobjects"
	// "backend/utils"

	// "code.google.com/p/go.crypto/bcrypt"
	// "github.com/hishboy/gocommons/lang"
)

func GetTurnsForMatch(matchId string, playerId string) ([]*models.Turn, *errors.ServerError) {
	turns, err := dao.GetTurns(matchId, playerId)
	if err != nil {
		return nil, errors.New(err, "Unable to retrieve turns", 500)
	}
	return turns, nil
}

func GetTurn(turnId string, playerId string) (*models.Turn, *errors.ServerError) {
	turn, err := dao.GetTurn(turnId, playerId)
	if err != nil {
		return nil, errors.New(err, "Unable to retrieve turn", 500)
	}
	return turn, nil
}

func MakeTurn(matchId string, playerId string, turnJson string) (*models.Turn, *errors.ServerError) {
	//get most recent turn
	lastTurn, lastTurnErr := dao.GetMostRecentTurn(matchId, playerId)
	if lastTurnErr != nil {
		return nil, errors.New(lastTurnErr, "Unable to retrieve last turn", 500)
	}

	playerIdInt, err := strconv.ParseInt(playerId, 10, 64)
	if err != nil {
		return nil, errors.New(err, "Unable to parse playerId", 422)
	}

	//check that player making turn didn't make last turn
	if lastTurn != nil && *lastTurn.PlayerId == playerIdInt {
		return nil, errors.New(nil, "Player made last turn", 400)
	}

	matchIdInt, err := strconv.ParseInt(matchId, 10, 64)
	if err != nil {
		return nil, errors.New(err, "Unable to parse matchId", 422)
	}

	// var jsonMap map[string]interface{}
	var jsonArray []map[string]interface{}
	jsonErr := json.Unmarshal([]byte(turnJson), &jsonArray)
	if jsonErr != nil {
		return nil, errors.New(jsonErr, "Unable to parse json", 500)
	}

	json := models.Json(jsonArray)

	var turnNumberInt int64
	turnNumberInt = 0
	if lastTurn != nil {
		turnNumberInt = *lastTurn.TurnNumber + 1
	}

	turn := &models.Turn{
		MatchId:    &matchIdInt,
		PlayerId:   &playerIdInt,
		TurnNumber: &turnNumberInt,
		Board:      &json,
	}
	var turnErr error

	turn, turnErr = dao.InsertTurn(turn)
	if turnErr != nil {
		return nil, errors.New(turnErr, "Unable to insert turn", 500)
	}

	return turn, nil
}
