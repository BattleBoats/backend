package handlers

import (
	// "errors"
	"fmt"
	"net/http"

	"backend/models"
	"backend/services"

	"github.com/codegangsta/martini"
)

type AppSession struct {
	Id string
}

type AppSessionPlayer struct {
	Id string
}

func Auth() martini.Handler {
	return func(c martini.Context, res http.ResponseWriter, req *http.Request) {
		// Check if app session cookie is present
		appSessionString := ""
		if cookie, err := req.Cookie(models.COOKIE_NAME); err == nil {
			appSessionString = cookie.Value
		} else {
			fmt.Println(err)
		}

		// fmt.Printf("appSessionString: %v\n", appSessionString)
		if appSessionString == "" {
			http.Error(res, "Not Authorized (appSessionString)", http.StatusUnauthorized)
			return
		}

		// Decode AppSession String into AppSession
		appSession := &AppSession{}
		models.AppCookie.Decode(models.COOKIE_NAME, appSessionString, appSession)

		// Validate Session
		validPlayerId := ""
		if appSession.Id != "" {
			// fmt.Println("ValidateSession id")
			playerId, err := services.ValidatePlayerSession(appSession.Id)
			if err != nil || len(playerId) == 0 {
				fmt.Println(err)
				http.Error(res, "Not Authorized (ValidateSession id)", http.StatusUnauthorized)
				return
			}
			validPlayerId = playerId
		}

		// fmt.Println("Valid user not found")
		if validPlayerId == "" {
			http.Error(res, "Not Authorized (Valid player not found)", http.StatusUnauthorized)
			return
		}

		// fmt.Println("c.Map()")
		c.Map((*AppSessionPlayer)(&AppSessionPlayer{validPlayerId}))
		c.Map((*AppSession)(appSession))
	}
}
