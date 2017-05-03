package handlers

import (
	"net/http"
)

func HandleAbout(w http.ResponseWriter, r *http.Request) {
	execTemplate("about", nil, r, w)
}
