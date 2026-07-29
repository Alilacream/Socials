package handlers

import (
	"context"
	"net/http"
	"time"

	"alilacream/socialx/internal/store"
	"alilacream/socialx/logs"
	"alilacream/socialx/models"
)

// Login function
func Login(store *store.UserStore) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := r.ParseForm(); err != nil {
			logs.DisplayErr(w, "Parse")
			return
		}
		LoggedUser := &models.User{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
		}
		ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
		defer cancel()
		// Checking the only the Login user if he exists
		if err := store.CheckUser(ctx, LoggedUser); err != nil {
			logs.DisplayErr(w, "Userlogin")
			return
		}
	}
}
