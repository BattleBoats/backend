package services

import (
	"fmt"
	// "log"
	// "time"
	"database/sql"
	"strconv"

	"backend/dao"
	"backend/errors"
	"backend/models"
	// "backend/models/sobjects"
	// "backend/utils"

	"code.google.com/p/go.crypto/bcrypt"
	// "github.com/huandu/facebook"
	// "github.com/nimajalali/go-force/force"
)

func GetUser(userId string) (*models.User, *errors.ServerError) {
	user := &models.User{}

	user, err := dao.GetUserById(userId)
	if err != nil {
		return nil, errors.New(err, "Unable to retrieve user", 500)
	}

	return user, nil
}

func LoginUser(email string, password string) (string, *models.User, *errors.ServerError) {

	if len(email) == 0 || len(password) == 0 {
		// utils.Count("User Login Failed Bad Request", 1)
		return "", nil, errors.New(nil, "Invalid username and password", 400)
	}

	//Get user
	user, err := dao.GetUserByEmail(email)
	if err != nil {
		return "", nil, errors.New(err, "Unable to retrieve user", 500)
	}

	if user == nil {
		return "", nil, errors.New(nil, "Not registered", 400)
	}

	// Check if valid password data exists
	if len(user.Password.String) == 0 {
		return "", nil, errors.New(nil, "Invalid User Object", 500)
	}

	// Validate password
	if passwordErr := bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(password)); passwordErr != nil {
		// utils.Count("User Login Failed Bad Password", 1)
		return "", nil, errors.New(passwordErr, "Unauthorized User", 401)
	}

	// Create Session
	sesh, _ := CreateUserSession(strconv.FormatInt(user.Id.Int64, 10))

	return sesh, user, nil
}

func RegisterUser(email, password string) (string, *models.User, *errors.ServerError) {
	// return "", nil, nil

	if len(password) == 0 {
		return "", nil, errors.New(nil, "User registration data incomplete", 400)
	}

	fmt.Println("Getting user by email")
	// Check if user exists by email
	user, serverErr := dao.GetUserByEmail(email)
	if serverErr != nil && serverErr != sql.ErrNoRows {
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

	fmt.Println("inserting user")
	user, err := dao.InsertUser(email, string(hashedPassword))
	if err != nil {
		return "", nil, errors.New(err, "Unable to insert user", 500)
	}

	// Create Session
	sesh, _ := CreateUserSession(strconv.FormatInt(user.Id.Int64, 10))

	return sesh, user, nil
}
