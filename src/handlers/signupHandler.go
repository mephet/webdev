package handlers

import (
	"net/http"
	"io/ioutil"
	"log"
	"github.com/gorilla/csrf"
	"html/template"
	"database/sql"
	_ "github.com/lib/pq"
	"../config"
)

func doPasswordsMatch(p1 string, p2 string) bool{
	if (p1 != p2) {
		return false
	} else {
		return true
	}
}

func doesEmailExist(email string) bool{
	var rowCount int

	db, err := sql.Open("postgres", config.DbInfo)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT count(*) as rowCount FROM UserAccounts WHERE Email = $1", email)
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		err = rows.Scan(&rowCount)
		if err != nil {
			log.Println(err)
		}
	}
	if rowCount == 0 {
		return false
	} else {
		return true
	}
}

func isInputPopulated(inputList []string) bool {
	for _, val := range inputList {
		if val == "" {
			return false
		}
	}
	return true
}

func addUserAccount(firstName string, lastName string, email string, password string, role string) bool{
	db, err := sql.Open("postgres", config.DbInfo)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	_, queryErr := db.Query("INSERT INTO useraccounts (email, password, role, firstname, lastname, accesslevel) VALUES($1, $2, $3, $4, $5, $6)", email, password, role, firstName, lastName, "STANDARD")
	if queryErr != nil {
		log.Println(queryErr)
		return false
	} else {
		log.Printf("User Account Created! Email: %s, Role: %s Inserted into UserAccounts Table.", email, role)
		return true
	}

}

func execTemplate(tName string, pipeString map[string]interface{}, r *http.Request, w http.ResponseWriter) {
	t := template.New(tName)
	path := r.URL.Path[1:];
	// Scan all files in dir static/templates and parse them into fileInfo
	dirName := config.TEMPLATE_DIR
	templateDir, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Println(err)
	}

	//Walk through the FileInfo array and Parse files as templates
	for _, templateName := range templateDir {
		t, _ = t.ParseFiles(dirName+"/"+templateName.Name()) // parsing of template string
	}
	t.ExecuteTemplate(w, path +".html", pipeString)
}


func HandleSignup(w http.ResponseWriter, r *http.Request) {

	var csrf = map[string]interface{}{
		csrf.TemplateTag: csrf.TemplateField(r),
	}
	execTemplate("signup", csrf, r, w)
}

func HandleSignupSubmit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	firstName := r.FormValue("firstname-input")
	lastName := r.FormValue("lastname-input")
	email := r.FormValue("email-input")
	role := r.FormValue("role-input")
	password := r.FormValue("password-input")
	passwordV := r.FormValue("password-verify")
	log.Println(firstName, lastName, email, role, password)

	if doesEmailExist(email) {
		m := make(map[string]interface{})
		m["email_error"] = "Email already exists!"
		m[csrf.TemplateTag] = csrf.TemplateField(r)
		execTemplate("signup", m, r, w)
	} else if !doPasswordsMatch(password, passwordV) {
		//Construct the template and pass the data interface (err + csrf) into the template on wrong password
		m := make(map[string]interface{})
		m["password_error"] = "Passwords do not match!"
		m[csrf.TemplateTag] = csrf.TemplateField(r)
		execTemplate("signup", m, r, w)
	} else {
		addUserAccount(firstName, lastName, email, password, role)
		redirectTarget := "/login"
		http.Redirect(w, r, redirectTarget, 302)
	}
}