package handlers

import (
	"net/http"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"database/sql"
)


func HandleIndex(w http.ResponseWriter, r *http.Request) {
	handleUserInput(w, r);
	t := template.New("main") //name of the template is main
	path := r.URL.Path[1:]

	// Scan all files in dir static/templates and parse them into fileInfo
	dirName := "static/templates"
	templateDir, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Println(err)
	}

	//Walk through the FileInfo array and Parse files as templates
	for _, templateName := range templateDir {
		t, _ = t.ParseFiles(dirName+"/"+templateName.Name()) // parsing of template string
	}
	t.ExecuteTemplate(w, path +".html", nil)
}

func HandleAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the about page")
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {

	t := template.New("login")
	path := r.URL.Path[1:]
	// Scan all files in dir static/templates and parse them into fileInfo
	dirName := "static/templates"
	templateDir, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Println(err)
	}

	//Walk through the FileInfo array and Parse files as templates
	for _, templateName := range templateDir {
		t, _ = t.ParseFiles(dirName+"/"+templateName.Name()) // parsing of template string
	}
	t.ExecuteTemplate(w, path +".html", nil)
}

func handleUserInput(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	usernameInput := r.FormValue("username-input")
	passwordInput := r.FormValue("password-input")
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_User, DB_Password, DB_Name)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	fmt.Println("# INSERTING VALUE...")
	rows , err := db.Query(`INSERT INTO userinfo(username, password, created) VALUES ($1, $2, $3)`, usernameInput, passwordInput, "2017-04-18")
	if err != nil {
		log.Println(err)
	} else {
		log.Println("User Successfully Inserted")
		log.Print(rows)
	}
}