package main

import (
	"fmt"
	"log"
	"net/http"

	"go-backend-starter/config"
	"go-backend-starter/handlers"
	"go-backend-starter/routes"
)

func main() {
	cfg := config.Load()

	userHandler := handlers.NewUserHandler()
	router := routes.NewRouter(userHandler)

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("server listening on http://localhost%s", addr)

	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
