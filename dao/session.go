package dao

import (
	"encoding/base32"
	"strings"
	"time"

	"backend/errors"
	"backend/models"

	"github.com/gorilla/securecookie"
)

var playerSessions map[string]models.PlayerSession

func init() {
	playerSessions = make(map[string]models.PlayerSession)
}

func GetPlayerSession(sessionId string) (*models.PlayerSession, *errors.ServerError) {
	playerSession, hasSession := playerSessions[sessionId]
	if hasSession != true {
		return nil, errors.New(nil, "Error on Get Player Session", 500)
	}

	if len(playerSession.Id) == 0 {
		return nil, nil
	}

	return &playerSession, nil
}

func SavePlayerSession(playerId string, expireTime time.Time) (string, *errors.ServerError) {
	playerSession := models.PlayerSession{
		PlayerId: playerId,
		Expire:   expireTime,
	}

	//New key is generated for each saved session, so a user could have multiple sessions.
	//Old sessions sit in the map, unless ExpireUserSession() is called.
	playerSession.Id = string(strings.TrimRight(
		base32.StdEncoding.EncodeToString(
			securecookie.GenerateRandomKey(32)), "="))

	playerSessions[playerSession.Id] = playerSession

	return playerSession.Id, nil
}

func ExpirePlayerSession(sessionId string) *errors.ServerError {
	_, sessionExists := playerSessions[sessionId]
	if sessionExists != true {
		return errors.New(nil, "Session does not exist", 500)
	}

	delete(playerSessions, sessionId)

	return nil
}
