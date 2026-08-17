package lib

import (
	"fmt"
	"net/http"
)

func DebugMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("=== REQUEST DEBUG ===")
		fmt.Printf("Authorization Header: %s\n", r.Header.Get("Authorization"))
		fmt.Printf("All Headers: %+v\n", r.Header)

		// Check cookie
		cookie, err := r.Cookie("jwt")
		if err == nil {
			fmt.Printf("JWT Cookie: %s\n", cookie.Value)
		} else {
			fmt.Println("No JWT cookie found")
		}
		fmt.Println("======================")

		next.ServeHTTP(w, r)
	})
}
