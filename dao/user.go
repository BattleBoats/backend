package dao

import (
	// "crypto/md5"
	"fmt"
	// "io"
	// "strings"
	// "time"

	// "backend/errors"
	"backend/models"
)

const (
	usersTable              = "users"
	userTagsTable           = "user_tags"
	userPasswordResetsTable = "user_password_resets"
)

func GetUser(userId string) (*models.User, error) {
	// session, serverErr := GetUserSession(sessionId)
	//    if serverErr != nil {
	//        return nil, serverErr
	//    }
	//    defer session.Close()

	//    var user models.DatabaseUserDoc

	//    err := r.Table(usersTable).Get(userId).Run(session).One(&user)
	//    if err != nil && err != (r.ErrWrongResponseType{}) {
	//        return nil, errors.New(err, "DB Error on Get User", 500)
	//    }

	//    if len(user.Id) == 0 {
	//        return nil, nil
	//    }

	//    return &user, nil
	return nil, nil
}

func GetUserByEmail(email string) (*models.User, error) {

	return nil, nil
}

func InsertUser(email string, password string) (*models.User, error) {
	nullEmail := models.BoxString(email)
	nullPassword := models.BoxString(password)

	fmt.Printf("email: %v\npassword: %v\n", nullEmail.String, nullPassword.String)

	user := &models.User{
		Email:    nullEmail,
		Password: nullPassword,
	}
	fmt.Printf("user: %v\n", user)

	dbMap, err := getDbMap()
	defer dbMap.Db.Close()
	if err != nil {
		return nil, err
	}
	fmt.Printf("dbmap: %v\n", dbMap)

	err = dbMap.Insert(user)
	if err != nil {
		return nil, err
	}
	fmt.Printf("user id: %v\n", user.Id)
	return user, nil
}
