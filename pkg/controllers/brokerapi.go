package controllers

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// Will contain the code that handles api endpoints

func (b *Base) RunBroker() {

	port := ":80"

	b.Router = chi.NewRouter()

	b.Router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"POST", "GET", "DELETE", "PUT", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Check if the broker works
	b.Router.Use(middleware.Heartbeat("/ping-broker"))

	log.Printf("Running broker service on port %s ... \n", port)
	srv := &http.Server{
		Addr:    port,
		Handler: b.Router,
	}

	// err := http.ListenAndServe(port, b.Router)

	err := srv.ListenAndServe()
	if err != nil {
		log.Printf("Error %s", err)
		os.Exit(1)
	}

}
