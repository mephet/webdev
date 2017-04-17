package controllers

import (
	"net/http"
	"fmt"
	"html/template"
)


func HandleIndex(w http.ResponseWriter, r *http.Request) {

	t := template.New("main") //name of the template is main
	t, _ = t.ParseFiles("static/templates/index.html", "static/templates/header.html" ) // parsing of template string
	t.ExecuteTemplate(w, "index.html", nil)
}

func HandleAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the about page")
}