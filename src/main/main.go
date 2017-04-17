package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"../controllers"
	"log"
)


var router = mux.NewRouter()



func main() {
	router.HandleFunc("/index", controllers.HandleIndex)
	router.HandleFunc("/about", controllers.HandleAbout)
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	http.Handle("/", router)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}

}

