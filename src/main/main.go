package main

import (
	"github.com/gorilla/mux"
	"log"
	"../handlers"
	"net/http"
)


var router = mux.NewRouter()

type Person struct {
	Name string
	Number string
}

//func getSession() *mgo.Session {
//	// connect to out local mgo
//	s, err := mgo.Dial("mongodb://localhost")
//	if err != nil {
//		panic(err)
//	}
//	return s
//}

func main() {
	//db := getSession()
	//defer db.Close()
	//c := db.DB("test").C("people")
	//err := c.Insert(&Person{"Ale", "+55 53 8116 9639"},
	//	&Person{"Cla", "+55 53 8402 8510"})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//result := Person{}
	//err = c.Find(bson.M{"name": "Ale"}).One(&result)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("Phone:", result.Number)

	router.HandleFunc("/index", handlers.HandleIndex)
	router.HandleFunc("/about", handlers.HandleAbout)
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

