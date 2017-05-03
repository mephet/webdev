package handlers

import (
	"net/http"
	"fmt"
)

func HandleAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the about page")
}
