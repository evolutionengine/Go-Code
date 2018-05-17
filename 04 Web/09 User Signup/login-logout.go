package main

import (
	"net/http"

	"github.com/gorilla/securecookie"
)

var cookieHandler = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

func getUserName(r *http.Request) (userName string) {
	cookie, err := r.Cookie("session")
	if err == nil {
		cookieValue := make(map[string]string)
		err = cookieHandler.Decode("session", cookie.Value, &cookieValue)
		if err == nil {
			userName = cookieValue["username"]
		}
	}
	return userName
}

func setCookie(w http.ResponseWriter, userName string) {

}
