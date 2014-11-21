package controllers

import (
	"net/http"

	"backend/handlers"

	"github.com/codegangsta/martini"
)

func init() {
	m.Post("/turns/:matchId", handlers.Auth(), handleMatchTurns)
	m.Post("/turn/:id", handlers.Auth(), handleTurn)
	m.Post("/turn/:matchId/make", handlers.Auth(), handleMakeTurn)
}

func handleMatchTurns(player *handlers.AppSessionPlayer, params martini.Params, r handlers.Respond, req *http.Request) {
	matchId := params["matchId"]
	turn, err := services.GetMatchTurns(matchId)
	if err != nil {
		r.Error(err)
		return
	}

	r.Valid(200, matchId)
}

func handleTurn(player *handlers.AppSessionPlayer, params martini.Params, r handlers.Respond, req *http.Request) {

}

func handleMakeTurn(player *handlers.AppSessionPlayer, params martini.Params, r handlers.Respond, req *http.Request) {

}
