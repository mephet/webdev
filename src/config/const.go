package config

import "fmt"

const DB_NAME = "WebDev"
const DB_USER = "postgres"
const DB_PASSWORD = "viewsonic1"

var DbInfo = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)

//Static Files
const TEMPLATE_DIR = "static/templates"

// Session Info
const DOMAIN = "localhost"
const SESSION_AGE = 30 * 60 //30 mins

