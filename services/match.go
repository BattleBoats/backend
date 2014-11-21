package services

import (
	// "fmt"
	// "log"
	// "time"
	// "database/sql"
	// "strconv"

	"backend/dao"
	"backend/errors"
	"backend/models"
	// "backend/models/sobjects"
	"backend/utils"

	// "code.google.com/p/go.crypto/bcrypt"
	// "github.com/hishboy/gocommons/lang"
)

var matchQueue *utils.MatchQueue

func init() {
	matchQueue = utils.NewQueue()
}

func GetMatches(playerId string, completed bool) ([]*models.Match, *errors.ServerError) {
	matches, err := dao.GetMatches(playerId, completed)
	if err != nil {
		return nil, errors.New(err, "Unable to retrieve matches", 500)
	}

	return matches, nil
}

func GetMatch(matchId string) (*models.Match, *errors.ServerError) {
	match, err := dao.GetMatchById(matchId)
	if err != nil {
		return nil, errors.New(err, "Unable to retrieve match", 500)
	}

	return match, nil
}

func FindMatch(playerId string) (*models.Match, *errors.ServerError) {
	//Look for an available Player
	player, err := GetPlayer(playerId)
	if err != nil {
		return nil, err
	}

	matchPlayer := matchQueue.PollInsert(player)

	if matchPlayer != nil {
		//set up match
		match, matchErr := dao.InsertMatch(*matchPlayer.Id, *player.Id)
		if matchErr != nil {
			return nil, errors.New(matchErr, "Could not create match", 500)
		}

		return match, nil
	}

	//no match has been created, but player is queued
	return nil, nil
}
