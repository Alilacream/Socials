package handlers

import (
	"encoding/json"
	"net/http"
)

// Welcome handler
func Welcome(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Hello and welcome to our server",
	})
}
