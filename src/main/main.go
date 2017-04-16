package main

import (
	"net/http"
	//"html/template"
	"github.com/gorilla/mux"
	"../controllers"
	"fmt"
)


var router = mux.NewRouter()

func init() {

}

func main() {
	fs := http.FileServer(http.Dir("static"))
	fmt.Println(fs)
	http.Handle("/css/", fs)
	router.HandleFunc("/", controllers.HandleIndex)
	router.HandleFunc("/about", controllers.HandleAbout)
	http.ListenAndServe(":8080", router)

}

