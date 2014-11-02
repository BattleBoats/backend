package dao

import (
	// "crypto/md5"
	// "fmt"
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

func GetUserById(userId string) (*models.User, error) {
	dbMap, err := getDbMap()
	defer dbMap.Db.Close()
	if err != nil {
		return nil, err
	}

	user := &models.User{}
	err = dbMap.SelectOne(&user, "select * from bb_user where bb_user_id=$1", userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	dbMap, err := getDbMap()
	defer dbMap.Db.Close()
	if err != nil {
		return nil, err
	}

	var user models.User
	err = dbMap.SelectOne(&user, "select * from bb_user where bb_user_email_address=$1", email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func InsertUser(email string, password string) (*models.User, error) {
	nullEmail := models.BoxString(email)
	nullPassword := models.BoxString(password)

	// fmt.Printf("email: %v\npassword: %v\n", nullEmail.String, nullPassword.String)

	user := &models.User{
		Email:    nullEmail,
		Password: nullPassword,
	}
	// fmt.Printf("user: %v\n", user)

	dbMap, err := getDbMap()
	defer dbMap.Db.Close()
	if err != nil {
		return nil, err
	}
	// fmt.Printf("dbmap: %v\n", dbMap)

	err = dbMap.Insert(user)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("user id: %v\n", user.Id)
	return user, nil
}
