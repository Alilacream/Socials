package handlers

import "net/http"

// Welcome handler
func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello and welcome to my Server"))
}
