package services

import (
	// "fmt"
	// "log"
	// "time"

	"backend/dao"
	"backend/errors"
	"backend/models"
	// "backend/models/sobjects"
	// "backend/utils"

	"code.google.com/p/go.crypto/bcrypt"
	// "github.com/huandu/facebook"
	// "github.com/nimajalali/go-force/force"
)

func RegisterUser(email, password string) (string, *models.User, *errors.ServerError) {
	// return "", nil, nil

	if len(password) == 0 {
		return "", nil, errors.New(nil, "User registration data incomplete", 400)
	}

	// Check if user exists by email
	user, serverErr := dao.GetUserByEmail(email)
	if serverErr != nil {
		return "", nil, errors.New(serverErr, "Unable to get user by email", 500)
	}

	if user != nil {
		return "", nil, errors.New(nil, "User already exists", 409)
	}

	// Create hashed password
	hashedPassword, passwordErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if passwordErr != nil {
		return "", nil, errors.New(passwordErr, "Unable to hash password", 500)
	}

	user, err := dao.InsertUser(email, string(hashedPassword))
	if err != nil {
		return "", nil, errors.New(err, "Unable to insert user", 500)
	}

	// Create Session
	sesh, _ := CreateUserSession(string(user.Id))

	return sesh, user, nil
}
