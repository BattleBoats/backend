package controllers

import (
	"fmt"
	"net/http"
	"time"

	"backend/errors"
	"backend/handlers"
	"backend/models"
	"backend/services"
	"backend/utils"

	// "github.com/codegangsta/martini"
)

const (
	kEmail    = "email"
	kPassword = "password"
	kHttp     = "http"
)

//UNITY3D only allows GET and POST, so all API calls will be through POST
//methods, including the appropriate HTTP verb as a parameter
func init() {
	m.Post("/player", handlers.Auth(), handlePlayer)
	m.Post("/player/register", handlePlayerRegister)
	m.Post("/player/login", handlePlayerLogin)
	m.Post("/player/logout", handlers.Auth(), handlePlayerLogout)
}

func handlePlayer(player *handlers.AppSessionPlayer, r handlers.Respond, req *http.Request) {
	fmt.Printf("playerId: %v", player.Id)
	dbUser, err := services.GetPlayer(player.Id)
	if err != nil {
		r.Error(err)
		return
	}

	r.Valid(200, dbUser)
}

func handlePlayerRegister(w http.ResponseWriter, req *http.Request, r handlers.Respond) {
	password := req.FormValue(kPassword)
	email := req.FormValue(kEmail)

	// Register User
	sesh, user, err := services.RegisterPlayer(email, password)
	if err != nil {
		r.Error(err)
		return
	}

	// User Registered and Logged in, set cookie
	err = setPlayerCookie(w, handlers.AppSession{Id: sesh})
	if err != nil {
		r.Error(err)
		return
	}

	r.Valid(200, user)
}

func handlePlayerLogin(r handlers.Respond, w http.ResponseWriter, req *http.Request) {
	//Login User
	sesh, _, err := services.LoginPlayer(req.FormValue(kEmail), req.FormValue(kPassword))
	if err != nil {
		r.Error(err)
		return
	}

	//User logged in, set cookie
	err = setPlayerCookie(w, handlers.AppSession{Id: sesh})
	if err != nil {
		r.Error(err)
		return
	}

	r.Valid(200, nil)
}

func handlePlayerLogout(session *handlers.AppSession, w http.ResponseWriter, r handlers.Respond) {
	// Invalidate session cookie
	cookie := &http.Cookie{
		Name:   models.COOKIE_NAME,
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)

	err := services.ExpirePlayerSession(session.Id)
	if err != nil {
		r.Error(err)
		return
	}

	r.Valid(200, nil)
}

func setPlayerCookie(w http.ResponseWriter, session handlers.AppSession) *errors.ServerError {
	// User logged in, set cookie
	encoded, err := models.AppCookie.Encode(models.COOKIE_NAME, session)
	if err != nil {
		return errors.New(err, "Unable to create cookie", 500)
	}

	secureBool := utils.Conf.GetBool("session.cookie.secure")

	cookie := &http.Cookie{
		Name:     models.COOKIE_NAME,
		Value:    encoded,
		Secure:   secureBool, //Set this to true when we have ssl
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Now().AddDate(50, 0, 0),
	}
	// fmt.Printf("\n\n cookie: %v\n\n", cookie)
	http.SetCookie(w, cookie)

	return nil
}
