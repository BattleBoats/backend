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

func CreateUserSession(userId string) (string, *errors.ServerError) {
	expireTime := time.Now().Add(time.Duration(sessionExpireTime) * time.Second)

	sessionString, err := dao.SaveUserSession(userId, expireTime)

	return sessionString, err
}

func ValidateUserSession(sessionId string) (string, *errors.ServerError) {
	userSession, err := dao.GetUserSession(sessionId)
	if err != nil {
		return "", err
	}

	if userSession == nil {
		return "", errors.New(nil, "Not Authenticated", 401)
	}

	// Verify not expired
	if userSession.Expire.Unix() < time.Now().Unix() {
		go ExpireUserSession(sessionId)
		return "", errors.New(nil, "Not Authenticated", 401)
	}

	return userSession.UserId, nil
}

func ExpireUserSession(sessionId string) *errors.ServerError {
	return dao.ExpireUserSession(sessionId)
}
