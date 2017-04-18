package models

import (
	"fmt"
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

const (
	DB_Name = "test"
	DB_User = "postgres"
	DB_Password = "viewsonic1"
)

func StartDB() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_User, DB_Password, DB_Name)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	fmt.Println("# INSERTING VALUE...")
	rows , err := db.Query(`INSERT INTO userinfo(username, departname, created) VALUES ('mr roth', 'mmo', '2012-09-09')`)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Inserted data: " + string(rows))
	}
}