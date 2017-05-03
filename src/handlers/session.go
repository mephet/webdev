package handlers

import (
	"github.com/gorilla/sessions"
	"encoding/gob"
	"net/http"
	"log"
)

type User struct {
	FirstName string
	LastName string
	Email string
	IsLoggedIn bool
}

var store = sessions.NewCookieStore([]byte("test"))

func init() {
	store.Options = &sessions.Options{
		Domain:   "localhost",
		Path:     "/",
		MaxAge:   60 * 5, // 5 mins
		HttpOnly: true,
	}
}

func CreateSession(w http.ResponseWriter, r *http.Request, firstName string, lastName string, email string) {
	session, err := store.Get(r, "loginSession")
	if err != nil {
		log.Println(err)
	}
	gob.Register(User{})
	user := User{firstName,lastName, email, true}
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
		user := session.Values["User"].(User)

		// Outputs user info to log
		log.Println(user)

		// If session expires, redirect to login page
		if !user.IsLoggedIn {
			http.Redirect(w, r, "/login", 301)
		}
	}


}