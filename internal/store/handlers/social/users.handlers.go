package social

import (
	"net/http"

	"alilacream/socialx/internal/store"
)

func FindUser(s *store.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
