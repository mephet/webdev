package main

import (
	"github.com/gorilla/mux"
	"log"
	"../handlers"
	"net/http"
	"github.com/gorilla/csrf"
)


var router = mux.NewRouter()


func main() {
	router.HandleFunc("/index", handlers.HandleIndex)
	router.HandleFunc("/about", handlers.HandleAbout)
	router.HandleFunc("/login", handlers.HandleLogin)
	router.HandleFunc("/signup", handlers.HandleSignup).Methods("GET")
	router.HandleFunc("/signup", handlers.HandleSignupSubmit).Methods("POST")

	// Fileserver to serve static files
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	http.Handle("/", router)
	err := http.ListenAndServe(":8080", csrf.Protect([]byte("32-byte-long-auth-key"), csrf.Secure(false))(router))
	if err != nil {
		log.Fatalln(err)
	}

}

