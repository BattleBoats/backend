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

//UNITY3D only allows GET and POST, so all API calls will be through POST
//methods, including the appropriate HTTP verb as a parameter
func init() {
	m.Post("/user", handlers.Auth(), handleUser)
	m.Post("/user/login", handleUserLogin)
	m.Post("/user/register", handleUserRegister)
}

func handleUser(user *handlers.AppSessionUser, r handlers.Respond, req *http.Request) {
	httpType := req.FormValue("http")
	r.Valid(200, httpType)
}

func handleUserLogin(r handlers.Respond, w http.ResponseWriter, req *http.Request) {
	//Login User
	// sesh, _, err := services.LoginUser(req.FormValue("email"), req.FormValue("password"))
	// if err != nil {
	// 	r.Error(err)
	// 	return
	// }

	// //User logged in, set cookie
	// err = setUserCookie(w, handlers.AppSession{Id: sesh})
	// if err != nil {
	// 	r.Error(err)
	// 	return
	// }

	r.Valid(200, nil)
}

func handleUserRegister(w http.ResponseWriter, req *http.Request, r handlers.Respond) {
	password := req.FormValue("password")
	email := req.FormValue("email")
	// Register User
	sesh, user, err := services.RegisterUser(email, password)
	if err != nil {
		r.Error(err)
		return
	}

	fmt.Printf("user: %v", user)

	// //create donor
	// donor := &models.Donor{
	// 	SoulSoupUser: user.Id,
	// }
	// fmt.Printf("\n\n user: %v \n\n", user)
	// resp, err := services.CreateDonor(donor)
	// if err != nil {
	// 	r.Error(err)
	// }

	// fmt.Printf("\n\n resp: %v \n\n", resp)

	// User Registered and Logged in, set cookie
	err = setUserCookie(w, handlers.AppSession{Id: sesh})
	if err != nil {
		r.Error(err)
		return
	}

	r.Valid(200, nil)
}

func setUserCookie(w http.ResponseWriter, session handlers.AppSession) *errors.ServerError {
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
