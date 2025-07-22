// Package config provides configuration utilities including server initialization.
package config

import (
	"log"
	"net/http"
	"os"
	"time"
)

// StartServer initializes and starts the HTTP server.
//
// It reads the port number from the "PORT" environment variable, then configures
// the server with reasonable defaults for write timeout, read timeout, and idle timeout.
//
// Parameters:
//   - router: the HTTP handler.
//
// Example usage:
//
//	router := NewRouter()  // assuming you've defined a router
//	config.StartServer(router)
//
// If the server fails to start, it logs the error and exits the application.
func StartServer(router http.Handler) {
	port := os.Getenv("PORT")
	srv := &http.Server{
		Handler:      router,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("Server starting on port " + port)
	log.Fatal(srv.ListenAndServe())
}
