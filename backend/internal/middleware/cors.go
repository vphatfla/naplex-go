package middleware

import "net/http"

func AddCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
   	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	next.ServeHTTP(w, r)
	})
}
