package handlers

import (
	"github.com/gorilla/sessions"
	"encoding/gob"
	"net/http"
	"log"
	"../config"
)

type User struct {
	FirstName string
	LastName string
	Email string
	AccessLevel string
	IsLoggedIn bool
}

var store = sessions.NewCookieStore([]byte("test"))

func init() {
	store.Options = &sessions.Options{
		Domain:   config.DOMAIN,
		Path:     "/",
		MaxAge:   config.SESSION_AGE,
		HttpOnly: true,
	}
}

func CreateSession(w http.ResponseWriter, r *http.Request, firstName string, lastName string, email string, accessLevel string) {
	session, err := store.Get(r, "loginSession")
	if err != nil {
		log.Println(err)
	}
	gob.Register(User{})
	user := User{firstName,lastName, email, accessLevel, true}
	session.Values["User"] = user
	err = session.Save(r, w)
	if err != nil {
		log.Println(err)
	}
}

func CheckValidLogin(w http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, "loginSession")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		// If session expires, redirect to login page
		if session.Values["User"] == nil {
			http.Redirect(w, r, "/login", 301)
		} else {
			// Refresh session Timer
			session.Options.MaxAge = config.SESSION_AGE
			user := session.Values["User"].(User)
			// Outputs user info to log
			log.Println("User login valid: " + user.Email)
		}
	}


}