package main

import (
	"log"
	"net/http"
	"time"

	"alilacream/socialx/internal/handler"
	"alilacream/socialx/internal/store"
	"alilacream/socialx/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type app struct {
	config Config
	store  store.Storage
}

type Config struct {
	addr string
	db   models.DBConfig
}

// setting up the new Server mux type with the routes within
func (a *app) route() *chi.Mux {
	mux := chi.NewMux()
	// good middleware stack
	// DOC: https://github.com/go-chi/chi
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.RequestID)
	// set timeout for response and request
	mux.Use(middleware.Timeout(time.Minute))
	// group all users in the v1, much more practical
	mux.Route("/v1", func(r chi.Router) {
		r.Get("/", handler.Welcome)
		//		r.Get("/users", handler.CreateAndInviteUsers)
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
