package social

import (
	"encoding/json"
	"log"
	"net/http"

	"alilacream/socialx/internal/store"
	"alilacream/socialx/models"

	"github.com/go-chi/chi/v5"
)

func FindUser(s *store.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		username := chi.URLParam(r, "username")
		if username == "" {
			http.Error(w, "User id is not a decimal", http.StatusBadRequest)
			return
		}
		user := models.User{Username: username}
		if err := s.Users.Search(r.Context(), &user); err != nil {
			log.Println("on why ?", err.Error())
			http.Error(w, "User Does not exist", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(map[string]models.User{
			"user has been found": user,
		})
	}
}
