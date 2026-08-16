package social

import (
	"encoding/json"
	"log"
	"net/http"

	"alilacream/socialx/internal/store"
	"alilacream/socialx/models"

	"github.com/go-chi/jwtauth/v5"
)

func Post(store *store.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var post models.Post

		w.Header().Set("Content-Type", "application/json")

		log.Println(r.Context())

		_, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		userID, ok := claims["sub"].(float64)
		if !ok || userID == 0 {
			log.Println("user id: ", userID)
			http.Error(w, "Invalid user ID", http.StatusUnauthorized)
			return
		}

		err = r.ParseForm()
		if err != nil || r.Body == nil {
			http.Error(w, "Couldn't Parse the request given", http.StatusBadRequest)
		}

		json.NewDecoder(r.Body).Decode(&post)
		// no empty fk refrence
		post.UserID = int64(userID)
		log.Println("POSTED: ", post)
		if len(post.Content) < 30 {
			http.Error(w, "The content needs to have at least 30 characters", http.StatusNotAcceptable)
			return
		}
		err = store.Posts.Create(r.Context(), &post)
		if err != nil {
			log.Println("here is why", err.Error())
			http.Error(w, "Couldn't store the post given", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(map[string]string{
			"success": "launched a new post",
		})
	}
}

func FindPost(s *store.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var post models.Post

		w.Header().Set("Content-Type", "application/json")
		err := r.ParseForm()
		if err != nil || r.Body == nil {
			http.Error(w, "Couldn't Parse the request given", http.StatusBadRequest)
		}
		json.NewDecoder(r.Body).Decode(&post)
		if err := s.Posts.Search(r.Context(), &post); err != nil {
			http.Error(w, "Couldn't find the post", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(map[string]models.Post{
			"Found it": post,
		})
	}
}

/*func Find_UserPosts(s *store.Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var post models.Post
		w.Header().Set("Content-Type", "application/json")
		err := r.ParseForm()
		if err != nil || r.Body == nil {
			http.Error(w, "Couldn't Parse the request given", http.StatusBadRequest)
		}
	}
}*/
