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

type AppSessionUser struct {
	Id string
}

func Auth() martini.Handler {
	return func(c martini.Context, res http.ResponseWriter, req *http.Request) {
		// fmt.Println(req)
		// Check if app session cookie is present
		appSessionString := ""
		if cookie, err := req.Cookie(models.COOKIE_NAME); err == nil {
			appSessionString = cookie.Value
			// fmt.Printf("cookie: %v\n", cookie)
		} else {
			fmt.Println(err)
			// fmt.Printf("cookie: %v\n", cookie)
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
		validUserId := ""
		if appSession.Id != "" {
			// fmt.Println("ValidateSession id")
			userId, err := services.ValidateUserSession(appSession.Id)
			if err != nil || len(userId) == 0 {
				fmt.Println(err)
				http.Error(res, "Not Authorized (ValidateSession id)", http.StatusUnauthorized)
				return
			}
			validUserId = userId
		}

		// fmt.Println("Valid user not found")
		if validUserId == "" {
			http.Error(res, "Not Authorized (Valid user not found)", http.StatusUnauthorized)
			return
		}

		// fmt.Println("c.Map()")
		c.Map((*AppSessionUser)(&AppSessionUser{validUserId}))
		c.Map((*AppSession)(appSession))
	}
}
