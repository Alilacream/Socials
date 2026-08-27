package handlers

import (
	"encoding/json"
	"net/http"

	"alilacream/socialx/internal/env"
	"alilacream/socialx/lib"
	"alilacream/socialx/models"
)

// Welcome handler
func Welcome(w http.ResponseWriter, r *http.Request) {
	userID, err := lib.GetUserIDFromCookie(env.GetVar("SECRET_KEY"), r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	user := models.User{ID: userID, Username: "testuser"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]*models.User{
		"user": &user,
	})
}
