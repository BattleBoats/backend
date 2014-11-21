package services

import (
	// "fmt"
	// "log"
	// "time"
	// "database/sql"
	// "strconv"

	"backend/dao"
	"backend/errors"
	"backend/models"
	// "backend/models/sobjects"
	"backend/utils"

	// "code.google.com/p/go.crypto/bcrypt"
	// "github.com/hishboy/gocommons/lang"
)

func GetMatchTurns(string matchId) ([]*models.Turn, *errors.ServerError) {
	var turns []*models.Turn
}
