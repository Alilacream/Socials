package main

import (
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/tomasen/realip"
	"golang.org/x/time/rate"
)

type client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

func (app *app) rateLimit(next *chi.Mux, rps, burst int) http.Handler {
	var (
		mu sync.Mutex

		clients = make(map[string]*client)
	)
	// freeing memory a little bit since a lot of concurrent users will leave
	go func() {
		time.Sleep(time.Minute) // runs every minute
		mu.Lock()
		defer mu.Unlock()
		for ip, c := range clients {
			if time.Since(c.lastSeen) > time.Minute*5 {
				delete(clients, ip)
			}
		}
	}()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		// getting the real ip
		ip := realip.FromRequest(r)
		if _, ok := clients[ip]; !ok {
			clients[ip] = &client{
				limiter:  rate.NewLimiter(rate.Limit(rps), burst),
				lastSeen: time.Now(),
			}
		}
		if !clients[ip].limiter.Allow() {
			app.appRateLimiterExceeded(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}
