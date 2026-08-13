package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"alilacream/socialx/internal/env"
	"alilacream/socialx/internal/store"
	"alilacream/socialx/logs"
	"alilacream/socialx/models"

	"github.com/golang-jwt/jwt/v5"
)

// Registering a new user with the help of UserStore Method (Create)
func Register(store *store.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var user models.User
		secret := env.GetVar("SECRET_KEY")
		// parsing the form to check if this format is goofy
		if err := r.ParseForm(); err != nil {
			logs.DisplayErr(w, "ParseForm")
			return
		}
		// if we couldn't process the format given
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Couldn't Process Request Body", http.StatusBadRequest)
			return
		}

		// using context
		// PERF:WHAT IS context ?
		ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
		defer cancel()
		// relativaly i should just passed the Users and call it a day, need to update the interfaces
		if err := store.Users.Create(ctx, &user); err != nil {
			log.Println("Error of creation: ", err.Error())
			http.Error(w, "User Already Exists: ", http.StatusNotAcceptable)
			return
		}
		// creating a map claim
		claims := jwt.MapClaims{
			"sub":                user.ID, // NOTE:even tho id was not grepped from the body.	it's value is scanned in the create user Method
			"preffered_username": user.Username,
			"email":              user.Email,
			"exp":                time.Now().Add(24 * time.Hour).Unix(),
			"iat":                time.Now().Unix(),
		}
		// Token Strucutre Creation
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenStr, err := token.SignedString([]byte(secret)) // Hmac expects an array of byte not a string
		if err != nil {

			log.Println("the Error of Auth", err.Error())
			http.Error(w, "Couldn't generate Token", http.StatusInternalServerError)
			return
		}
		// set it in the browser
		http.SetCookie(w, &http.Cookie{
			Name:     "jwt_token",
			Value:    tokenStr,
			HttpOnly: true,
		})
		// for debugging
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "New User Registered"})
	}
}
