package services

import (
	"time"

	"backend/dao"
	"backend/errors"
	"backend/utils"
)

var sessionExpireTime int

func init() {
	sessionExpireTime = utils.Conf.GetInt("session.expiretime")
}

func CreatePlayerSession(playerId string) (string, *errors.ServerError) {
	expireTime := time.Now().Add(time.Duration(sessionExpireTime) * time.Second)

	sessionString, err := dao.SavePlayerSession(playerId, expireTime)

	return sessionString, err
}

func ValidatePlayerSession(sessionId string) (string, *errors.ServerError) {
	playerSession, err := dao.GetPlayerSession(sessionId)
	if err != nil {
		return "", err
	}

	if playerSession == nil {
		return "", errors.New(nil, "Not Authenticated", 401)
	}

	// Verify not expired
	if playerSession.Expire.Unix() < time.Now().Unix() {
		go ExpirePlayerSession(sessionId)
		return "", errors.New(nil, "Not Authenticated", 401)
	}

	return playerSession.PlayerId, nil
}

func ExpirePlayerSession(sessionId string) *errors.ServerError {
	return dao.ExpirePlayerSession(sessionId)
}
