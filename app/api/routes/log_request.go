package routes

import (
	"fmt"
	"net/http"
)

func LogRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the request method and path
		fmt.Printf("Request: %s %s\n", r.Method, r.URL.Path)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
