package config

import "fmt"

const DB_NAME = "postgres"
const DB_USER = "postgres"
const DB_PASSWORD = "viewsonic1"

var DbInfo = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)



