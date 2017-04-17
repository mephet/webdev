package handlers

import (
	"net/http"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
)


func HandleIndex(w http.ResponseWriter, r *http.Request) {

	t := template.New("main") //name of the template is main
	path := r.URL.Path[1:]
	log.Println(path)

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