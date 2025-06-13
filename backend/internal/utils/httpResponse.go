package utils

import (
	"encoding/json"
	"net/http"
)

func HTTPJsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	// Needed for logging middleware, without explic WriteHeader, the middleware won't be able to get the status code
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(data)
}

func HTTPJsonError(w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{
		"error": msg,
	})
}
