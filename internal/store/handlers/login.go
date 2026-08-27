package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"alilacream/socialx/internal/store"
	"alilacream/socialx/lib"
	"alilacream/socialx/models"
)

// Login function
func Login(store *store.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var user models.User
		err := r.ParseForm()
		if err != nil || r.Body == nil {
			http.Error(w, "Couldn't parse the body Request Sent", http.StatusBadRequest)
			return
		}

		json.NewDecoder(r.Body).Decode(&user)
		ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
		defer cancel()
		// Checking the only the Login user if he exists
		err = store.Users.Check_User_Exist(ctx, &user)
		if err != nil {
			log.Println("[ERROR]:", err.Error())
			http.Error(w, "User Does not Exist", http.StatusNotFound)
			return
		}

		// generating the string token user
		tokenStr, err := lib.GenerateJWT(store.JWTSecret, &user)
		if err != nil {
			http.Error(w, "Couldn't generate Token", http.StatusInternalServerError)
		}

		// save the jwt token as a cookie.
		http.SetCookie(w, &http.Cookie{
			Name:     "jwt",
			Value:    tokenStr,
			HttpOnly: true,
			Path:     "/api", // <- determines where the cookie is sent, / for all
			MaxAge:   86400,
		})
		// all good
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(map[string]*models.User{
			"user": &user,
		})
	}
}
