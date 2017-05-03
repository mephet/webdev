package handlers

import (
	"net/http"
	"io/ioutil"
	"log"
	"html/template"
	"github.com/gorilla/csrf"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	CreateSession(w, r, "Dave", "Yogpod", "dave@hp.com")
	t := template.New("login")
	path := r.URL.Path[1:]
	// Scan all files in dir static/templates and parse them into fileInfo
	dirName := "static/templates"
	templateDir, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Print("Template cant be read")
		log.Println(err)
	}
	//
	////Walk through the FileInfo array and Parse files as templates
	for _, templateName := range templateDir {
		t, err = t.ParseFiles(dirName + "/" + templateName.Name()) // parsing of template string
		if err != nil {
			log.Println(err)
		}
	}
	var csrf = map[string]interface{}{
		csrf.TemplateTag: csrf.TemplateField(r),
	}
	t.ExecuteTemplate(w, path+".html", csrf)
}