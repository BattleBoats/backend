package services

import (
	// "fmt"
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
)

func GetPlayer(playerId string) (*models.Player, *errors.ServerError) {
	player, err := dao.GetPlayerById(playerId)
	if err != nil {
		return nil, errors.New(err, "Unable to retrieve player", 500)
	}

	return player, nil
}

func LoginPlayer(email string, password string) (string, *models.Player, *errors.ServerError) {

	if len(email) == 0 || len(password) == 0 {
		// utils.Count("User Login Failed Bad Request", 1)
		return "", nil, errors.New(nil, "Invalid username and password", 400)
	}

	//Get user
	player, err := dao.GetPlayerByEmail(email)
	if err != nil {
		return "", nil, errors.New(err, "Unable to retrieve user", 500)
	}

	if player == nil {
		return "", nil, errors.New(nil, "Not registered", 400)
	}

	// Check if valid password data exists
	if player.Password == nil || len(*player.Password) == 0 {
		return "", nil, errors.New(nil, "Invalid User Object", 500)
	}

	// Validate password
	if passwordErr := bcrypt.CompareHashAndPassword([]byte(*player.Password), []byte(password)); passwordErr != nil {
		// utils.Count("User Login Failed Bad Password", 1)
		return "", nil, errors.New(passwordErr, "Unauthorized player", 401)
	}

	// Create Session
	sesh, _ := CreatePlayerSession(strconv.FormatInt(*player.Id, 10))

	return sesh, player, nil
}

func RegisterPlayer(email, password string) (string, *models.Player, *errors.ServerError) {
	// return "", nil, nil

	if len(password) == 0 {
		return "", nil, errors.New(nil, "Player registration data incomplete", 400)
	}

	// Check if user exists by email
	player, serverErr := dao.GetPlayerByEmail(email)
	if serverErr != nil && serverErr != sql.ErrNoRows {
		return "", nil, errors.New(serverErr, "Unable to get player by email", 500)
	}

	if player != nil {
		return "", nil, errors.New(nil, "Player already exists", 409)
	}

	// Create hashed password
	hashedPassword, passwordErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if passwordErr != nil {
		return "", nil, errors.New(passwordErr, "Unable to hash password", 500)
	}

	player, err := dao.InsertPlayer(email, string(hashedPassword))
	if err != nil {
		return "", nil, errors.New(err, "Unable to insert player", 500)
	}

	// Create Session
	sesh, _ := CreatePlayerSession(strconv.FormatInt(*player.Id, 10))

	return sesh, player, nil
}
