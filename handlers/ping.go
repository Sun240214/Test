package handlers

import "net/http"

// Ping returns a small health-check response.
func Ping(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"message": "pong"})
}
