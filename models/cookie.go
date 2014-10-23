package models

import (
	"github.com/gorilla/securecookie"
)

const (
	COOKIE_NAME string = "battleboats-session"
)

var hashKey = []byte("VRa7fRQcnq7bXEDK4RO7lLIwBRvP2KtP")
var blockKey = []byte("RjLyp4V22g0qp1KmVgJDqSTlmzSbRFff")

var AppCookie = securecookie.New(hashKey, blockKey)
