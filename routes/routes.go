package routes

import (
	"net/http"

	"go-backend-starter/handlers"
)

// NewRouter registers API routes and shared middleware.
func NewRouter(userHandler *handlers.UserHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /ping", handlers.Ping)
	mux.HandleFunc("GET /users", userHandler.List)
	mux.HandleFunc("POST /users", userHandler.Create)

	return Chain(mux, Recover, Logging)
}
