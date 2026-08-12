package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"alilacream/socialx/internal/env"
	"alilacream/socialx/internal/store"
	"alilacream/socialx/logs"
	"alilacream/socialx/models"

	"github.com/golang-jwt/jwt/v5"
)

// Login function
func Login(store *store.Storage) func(w http.ResponseWriter, r *http.Request) {
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
		User, err := store.Users.Check(ctx, LoggedUser)
		if err != nil {
			logs.DisplayErr(w, "Userlogin")
			return
		}

		// the format is containing the payload, the header is the first param, containing the method
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub":       User.ID, // Use id, to be unique
			"firstname": User.FirstName,
			"lastname":  User.LastName,
			"user":      User.Username,
			"exp":       time.Now().Add(24 * time.Hour).Unix(),
			"iat":       time.Now().Unix(),
		})
		// the token is of type token, we need it's string value
		tokenStr, err := token.SignedString(env.GetVar("SECRET_KEY"))
		if err != nil {
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
		// last thing to show
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"server": "Welcome Back " + User.Username,
		})
	}
}
