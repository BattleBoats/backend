package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"backend/errors"
	"backend/handlers"
	"backend/services"

	"github.com/codegangsta/martini"
)

func init() {
	m.Post("/turns/:matchId", handlers.Auth(), handleMatchTurns)
	m.Post("/turn/:matchId/make", handlers.Auth(), handleMakeTurn)
	m.Post("/turn/:id", handlers.Auth(), handleTurn)
}

func handleMatchTurns(player *handlers.AppSessionPlayer, params martini.Params, r handlers.Respond, req *http.Request) {
	matchId := params["matchId"]
	turns, err := services.GetTurnsForMatch(matchId, player.Id)
	if err != nil {
		r.Error(err)
		return
	}

	r.Valid(200, turns)
}

func handleTurn(player *handlers.AppSessionPlayer, params martini.Params, r handlers.Respond, req *http.Request) {
	turnId := params["id"]
	turn, err := services.GetTurn(turnId, player.Id)
	if err != nil {
		r.Error(err)
		return
	}

	r.Valid(200, turn)
}

func handleMakeTurn(player *handlers.AppSessionPlayer, params martini.Params, r handlers.Respond, req *http.Request) {
	matchId := params["matchId"]

	bodyBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		r.Error(errors.New(err, "Unable to read body bytes", 400))
		return
	}

	fmt.Printf("body bytes: %v\n\n", string(bodyBytes))

	turn, insertErr := services.MakeTurn(matchId, player.Id, string(bodyBytes))
	if insertErr != nil {
		fmt.Println("Error inserting turn")
		r.Error(insertErr)
		return
	}

	r.Valid(200, turn)
}
