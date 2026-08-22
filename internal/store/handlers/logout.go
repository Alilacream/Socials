package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"alilacream/socialx/internal/store"
)

func Logout(s *store.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:     "jwt",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour), // expires immeadiatly
			HttpOnly: true,
			MaxAge:   -1, // deleted immeadiatly
			Path:     "/api",
		})
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"success": "you've logged out",
		})
	}
}
