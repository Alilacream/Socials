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

// Register  new user with the help of UserStore Method (Create)
func Register(store *store.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var user models.User
		// parsing the form to check if this format is goofy
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Bad json request format", http.StatusBadRequest)
			return
		}
		json.NewDecoder(r.Body).Decode(&user)
		// if the user input is invalid we return an http error
		if err := lib.ParseEmail(user.Email); err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}

		// if the email input is invalid we return an http error
		// HACK: need to add the @... testcase with valid domains ofc
		if err := lib.ParseUsername(user.Username); err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}

		// PERF:WHAT IS context ?
		ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
		defer cancel()
		// relativaly i should just passed the Users and call it a day, need to update the interfaces
		if err := store.Users.Create(ctx, &user); err != nil {
			log.Println("Error of creation: ", err.Error())
			http.Error(w, "User Already Exists: ", http.StatusNotAcceptable)
			return
		}
		tokenStr, err := lib.GenerateJWT(store.JWTSecret, &user)
		if err != nil {
			log.Println("the Error of Auth", err.Error())
			http.Error(w, "Couldn't generate Token", http.StatusInternalServerError)
			return
		}
		// set it in the browser
		http.SetCookie(w, &http.Cookie{
			Name:     "jwt",
			Value:    tokenStr,
			HttpOnly: true,
			Path:     "/api",
			MaxAge:   86400,
		})

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "New User Registered"})
	}
}
