package handlers

import (
	"net/http"
	"log"
	"github.com/gorilla/csrf"
	_ "github.com/lib/pq"
	"database/sql"
	"../config"
)


func HandleLogin(w http.ResponseWriter, r *http.Request) {

	var csrf = map[string]interface{}{
		csrf.TemplateTag: csrf.TemplateField(r),
	}
	execTemplate("login", csrf, r, w)
}

func HandleLoginSubmit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.FormValue("email-input")
	password := r.FormValue("password-input")

	isValid, firstName, lastName, accessLevel := validateLogin(email, password)
	if isValid {
		CreateSession(w, r, firstName, lastName, email, accessLevel)
		http.Redirect(w, r, "/index", 301)
	} else {
		m := make(map[string]interface{})
		m["loginError"] = "Incorrect Email/Password Combination!"
		m[csrf.TemplateTag] = csrf.TemplateField(r)
		execTemplate("login", m, r, w)
	}
}


func validateLogin(email string, password string) (bool, string, string, string) {
	db, err := sql.Open("postgres", config.DbInfo)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	var emailResult string
	var firstName string
	var lastName string
	var accessLevel string
	queryErr := db.QueryRow("SELECT email, firstname, lastname, accesslevel FROM useraccounts WHERE email = $1 AND password = $2", email, password).Scan(&emailResult, &firstName, &lastName, &accessLevel)
	switch {
	case queryErr == sql.ErrNoRows:
		log.Println("Incorrect Username/Password combination")
		return false, firstName, lastName, accessLevel
	case queryErr != nil:
		log.Println(queryErr)
		return false, firstName, lastName, accessLevel
	default:
		return true, firstName, lastName, accessLevel
	}
}