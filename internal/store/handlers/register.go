package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"alilacream/socialx/internal/env"
	"alilacream/socialx/internal/store"
	"alilacream/socialx/lib"
	"alilacream/socialx/logs"
	"alilacream/socialx/models"

	"github.com/golang-jwt/jwt/v5"
)

// Registering a new user with the help of UserStore Method (Create)
func Register(store *store.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

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

		formValues := &models.User{
			FirstName: r.FormValue("first_name"),
			LastName:  r.FormValue("last_name"),
			Username:  r.FormValue("username"),
			Email:     r.FormValue("email"),
			Password:  hashPass,
		}

		log.Println("value:", formValues)
		// using context
		// PERF:WHAT IS context ?
		ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
		defer cancel()
		// relativaly i should just passed the Users and call it a day, need to update the interfaces
		if err := store.Users.Create(ctx, formValues); err != nil {
			logs.Errors(w, "Userregister")
			return
		}
		// creating a map claim
		claims := jwt.MapClaims{
			"sub":  formValues.ID, // NOTE:even tho id was not grepped from the body.	it's value is scanned in the create user Method
			"user": formValues.Username,
			"exp":  time.Now().Add(24 * time.Hour).Unix(),
			"iat":  time.Now().Unix(),
		}
		// Token Strucutre Creation
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenStr, err := token.SignedString(env.GetVar("SECRET_KEY"))
		if err != nil {
			http.Error(w, "Couldn't generate Token", http.StatusInternalServerError)
			return
		}
		// set it in the browser
		http.SetCookie(w, &http.Cookie{
			Name:     "jwt_token",
			Value:    tokenStr,
			HttpOnly: true,
			Secure:   false, // we're only in local dev

		})
		// for debugging
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "New User Registered"})
	}
}
