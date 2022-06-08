package handlers

import (
	"encoding/json"
	"net/http"
)

func Routes() {
	http.HandleFunc("/sendJSON", SendJSON)
}

func SendJSON(w http.ResponseWriter, r *http.Request) {
	u := struct {
		Name  string
		Email string
	}{
		Name:  "Bill",
		Email: "bill@clifton.org",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&u)
}