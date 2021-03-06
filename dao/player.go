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
	kPLAYER_TABLE         = "bb_player"
	kPLAYER_ID            = "player_id"
	kPLAYER_EMAIL_ADDRESS = "email_address"
)

func GetPlayerById(playerId string) (*models.Player, error) {
	dbMap, err := getDbMap()
	if err != nil {
		return nil, err
	}
	defer dbMap.Db.Close()

	player := &models.Player{}
	query := fmt.Sprintf("SELECT * FROM %v WHERE %v=$1", kPLAYER_TABLE, kPLAYER_ID)
	err = dbMap.SelectOne(&player, query, playerId)
	if err != nil {
		return nil, err
	}

	return player, nil
}

func GetPlayerByEmail(email string) (*models.Player, error) {
	dbMap, err := getDbMap()
	if err != nil {
		return nil, err
	}
	defer dbMap.Db.Close()

	var player models.Player
	query := fmt.Sprintf("SELECT * FROM %v WHERE %v=$1", kPLAYER_TABLE, kPLAYER_EMAIL_ADDRESS)
	err = dbMap.SelectOne(&player, query, email)
	if err != nil {
		return nil, err
	}

	return &player, nil
}

func InsertPlayer(email string, password string) (*models.Player, error) {

	// fmt.Printf("email: %v\npassword: %v\n", nullEmail.String, nullPassword.String)

	player := &models.Player{
		Email:    &email,
		Password: &password,
	}
	// fmt.Printf("user: %v\n", user)

	dbMap, err := getDbMap()
	if err != nil {
		return nil, err
	}
	defer dbMap.Db.Close()
	// fmt.Printf("dbmap: %v\n", dbMap)

	err = dbMap.Insert(player)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("user id: %v\n", user.Id)
	return player, nil
}
