package middleware

import (
	"log"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	code int
	body []byte
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *responseWriter) Write(b []byte) (int, error) {
	w.body = append(w.body, b...)
	return w.ResponseWriter.Write(b)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapped := &responseWriter{
			ResponseWriter: w,
			code:           0,
		}
		next.ServeHTTP(wrapped, r)
		log.Printf(
			"[%s] %s %s | Status: %d | Duration: %v | Response: %s",
			time.Now().Format("2006-01-02 15:04:05"),
			r.Method,
			r.URL.Path,
			wrapped.code,
			time.Since(start),
			string(wrapped.body),
		)
	})
}
