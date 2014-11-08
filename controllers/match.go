package controllers

import (
	// "fmt"
	"net/http"
	"strconv"
	// "time"

	"backend/errors"
	"backend/handlers"
	// "backend/models"
	"backend/services"
	// "backend/utils"

	"github.com/codegangsta/martini"
)

func init() {
	m.Post("/matches/:complete", handlers.Auth(), handleAllMatches)
	m.Post("/match/find", handlers.Auth(), handleMatchFind)
	m.Post("/match/:id", handlers.Auth(), handleMatch)
}

func handleAllMatches(player *handlers.AppSessionPlayer, params martini.Params, r handlers.Respond, req *http.Request) {
	completedString := params["complete"]
	completed, err := strconv.ParseBool(completedString)
	if err != nil {
		r.Error(errors.New(err, "Unable to parse bool url parameter", 400))
		return
	}

	matches, matchesErr := services.GetMatches(player.Id, completed)
	if err != nil {
		r.Error(matchesErr)
		return
	}

	r.Valid(200, matches)
}

func handleMatch(player *handlers.AppSessionPlayer, params martini.Params, r handlers.Respond, req *http.Request) {
	// update getmatch to only look for player's match
	match, err := services.GetMatch(params["id"])
	if err != nil {
		r.Error(err)
		return
	}

	r.Valid(200, match)
}

/*
 * Look for a match with only one player in it.
 * If no open matches exist, create a new match with this player as player_one.
 */
func handleMatchFind(player *handlers.AppSessionPlayer, r handlers.Respond, req *http.Request) {
	match, err := services.FindMatch(player.Id)
	if err != nil {
		r.Error(err)
		return
	}

	if match == nil {
		r.Valid(202, nil)
		return
	}

	r.Valid(200, match)
}
