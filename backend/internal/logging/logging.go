package logging

import (
	"log"
	"net/http"
	"time"
)


type responseWriter struct{
	http.ResponseWriter
	code int
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrapped := &responseWriter {
			ResponseWriter: w,
			code: 0,
		}
		next.ServeHTTP(wrapped, r)
		log.Printf("StatusCode %v  |   Request URL %v   |   At %v", wrapped.code, r.URL.Path, time.Now())
	})
}
