package handler

import "net/http"

func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello and welcome to my Server"))
}
