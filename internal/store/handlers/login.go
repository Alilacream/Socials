package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"alilacream/socialx/internal/env"
	"alilacream/socialx/internal/store"
	"alilacream/socialx/models"

	"github.com/golang-jwt/jwt/v5"
)

// Login function
func Login(store *store.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		secret := env.GetVar("SECRET_KEY")

		var user models.User
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Couldn't parse the body Request Sent", http.StatusBadRequest)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Couldn't Process the Json Format", http.StatusInternalServerError)
		}

		ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
		defer cancel()
		// Checking the only the Login user if he exists
		err := store.Users.Check_User_Exist(ctx, &user)
		if err != nil {
			http.Error(w, "User Does not Exist", http.StatusNotFound)
			return
		}

		// the format is containing the payload, the header is the first param, containing the method
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub":                user.ID, // Use id, to be unique
			"preferred_username": user.Username,
			"email":              user.Email,
			"exp":                time.Now().Add(24 * time.Hour).Unix(),
			"iat":                time.Now().Unix(),
		})
		// the token is of type token, we need it's string value, NOTE: Hmac Meth needs the secret in a slice byte type
		tokenStr, err := token.SignedString([]byte(secret))
		if err != nil {
			log.Println("the Error of Auth", err.Error())
			http.Error(w, "Could not create token", http.StatusInternalServerError)
			return
		}
		// save the jwt token as a cookie.
		http.SetCookie(w, &http.Cookie{
			Name:     "jwt_token",
			Value:    tokenStr,
			HttpOnly: true,
			Secure:   false, // we're only in local dev
		})
		// all good
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"server": "Welcome Back " + user.Username,
		})
	}
}
