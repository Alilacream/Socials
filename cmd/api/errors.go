package main

import "net/http"

type HTTPError interface {
	appRateLimiterExceeded(w http.ResponseWriter, r *http.Request)
}

func (a *app) appRateLimiterExceeded(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "rate limite exceeded", http.StatusTooManyRequests)
}
