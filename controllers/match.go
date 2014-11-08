package controllers

import (
	// "fmt"
	"net/http"
	// "time"

	// "backend/errors"
	"backend/handlers"
	// "backend/models"
	"backend/services"
	// "backend/utils"

	"github.com/codegangsta/martini"
)

func init() {
	m.Post("/matches", handlers.Auth(), handleAllMatches)
	m.Post("/match/find", handlers.Auth(), handleMatchFind)
	m.Post("/match/:id", handlers.Auth(), handleMatch)
}

func handleAllMatches(r handlers.Respond, req *http.Request) {

	r.Valid(200, nil)
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
