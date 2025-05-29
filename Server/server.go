package server

import (
	"encoding/json"
	"net/http"

	files "github.com/anicse37/RestFullAPI/Files"
)

func GetUserData(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	file := files.Setup()

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(file.Users)

}
