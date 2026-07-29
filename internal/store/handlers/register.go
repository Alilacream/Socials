package handler

import (
	"context"
	"log"
	"net/http"
	"time"

	"alilacream/socialx/internal/store"
	"alilacream/socialx/lib"
	"alilacream/socialx/logs"
	"alilacream/socialx/models"
)

// Welcome handler
func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello and welcome to my Server"))
}

// Registering
func Register(w http.ResponseWriter, r *http.Request, store *store.UserStore) {
	w.Header().Set("Content-Type", "application/json")
	var formValues *models.User
	if err := r.ParseForm(); err != nil {
		logs.DisplayErr(w, "ParseForm")
		return
	}
	// HACK: need to hash ofc
	hashPass, err := lib.HashPassword(r.FormValue("password"))
	if err != nil {
		logs.DisplayErr(w, "BadRequest")
		return
	}
	formValues = &models.User{
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
		Username:  r.FormValue("username"),
		Password:  hashPass,
	}
	// using context
	ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
	defer cancel()
	if err := store.Create(ctx, formValues); err != nil {
		logs.Errors(w, "Userregister")
		return
	}

	log.Println("New User created: ", formValues.ID)
}
