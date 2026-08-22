package main

import (
	"log"
	"net/http"
	"time"

	"alilacream/socialx/internal/store"
	"alilacream/socialx/internal/store/handlers"
	"alilacream/socialx/internal/store/handlers/social"
	"alilacream/socialx/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
)

type app struct {
	config Config
	store  *store.Storage
}

type Config struct {
	addr string
	db   *models.DBConfig
}

// setting up the new Server mux type with the routes within
func (a *app) route(s *store.Storage) *chi.Mux {
	tokenAuth := jwtauth.New("HS256", []byte(s.JWTSecret), nil)
	mux := chi.NewMux()
	// good middleware stackPlain
	// DOC: https://github.com/go-chi/chi
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.RequestID)
	// set timeout for response and request
	mux.Use(middleware.Timeout(time.Minute))
	// protected routes, verifying in the jwt is valid or not

	// Source - https://stackoverflow.com/a/57682227
	// Posted by Santiago, modified by community. See post 'Timeline' for change history
	// Retrieved 2026-08-22, License - CC BY-SA 4.0

	mux.Route("/api", func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))

		r.Get("/", handlers.Welcome)
		r.Post("/post", social.Post(s))
		r.Get("/posts/{postID}", social.FindPost(s))
		r.Get("/users/{username}", social.FindUser(s))
		r.Get("/users/{username}/posts", social.Find_UserPosts(s))
	})

	// pub routes, group all users in the v1, much more practical
	mux.Route("/v1", func(r chi.Router) {
		r.Post("/register", handlers.Register(s))
		r.Post("/login", handlers.Login(s))
		r.Post("/logout", handlers.Logout(s))
	})
	return mux
}

// running the application, core method for serving
func (a *app) run(mux *chi.Mux) error {
	srv := &http.Server{
		Addr:    a.config.addr,
		Handler: mux,
		// security enhancing args in server interface
		WriteTimeout: time.Second * 30, // max timeout to write response to the client
		ReadTimeout:  time.Second * 10, // max timeout to read the request from the client
		IdleTimeout:  time.Minute,      // keeps the connection alive only in one minute
	}
	log.Println("Server listening in port: ", a.config.addr)
	return srv.ListenAndServe()
}
