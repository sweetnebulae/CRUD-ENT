package config

import (
	"log"
	"net/http"
	"os"
	"time"
)

func StartServer(router http.Handler) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	srv := &http.Server{
		Handler:      router,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("Server starting on port" + port)
	log.Fatal(srv.ListenAndServe())
}
