package main

import (
	"github.com/gorilla/mux"
	"log"
	"../handlers"
	"net/http"
)


var router = mux.NewRouter()


func main() {
	router.HandleFunc("/index", handlers.HandleIndex)
	router.HandleFunc("/about", handlers.HandleAbout)
	router.HandleFunc("/login", handlers.HandleLogin)
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

