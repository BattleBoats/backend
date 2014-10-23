package dao

import (
	"encoding/base32"
	"strings"
	"time"

	"backend/errors"
	"backend/models"

	"github.com/gorilla/securecookie"
)

var userSessions map[string]models.UserSession

func init() {
	userSessions = make(map[string]models.UserSession)
}

func GetUserSession(sessionId string) (*models.UserSession, *errors.ServerError) {
	userSession, hasSession := userSessions[sessionId]
	if hasSession != true {
		return nil, errors.New(nil, "Error on Get User Session", 500)
	}

	if len(userSession.Id) == 0 {
		return nil, nil
	}

	return &userSession, nil
}

func SaveUserSession(userId string, expireTime time.Time) (string, *errors.ServerError) {
	userSession := models.UserSession{
		UserId: userId,
		Expire: expireTime,
	}

	//New key is generated for each saved session, so a user could have multiple sessions.
	//Old sessions sit in the map, unless ExpireUserSession() is called.
	userSession.Id = string(strings.TrimRight(
		base32.StdEncoding.EncodeToString(
			securecookie.GenerateRandomKey(32)), "="))

	userSessions[userSession.Id] = userSession

	return userSession.Id, nil
}

func ExpireUserSession(sessionId string) *errors.ServerError {
	_, sessionExists := userSessions[sessionId]
	if sessionExists != true {
		return errors.New(nil, "Session does not exist", 500)
	}

	delete(userSessions, sessionId)

	return nil
}
