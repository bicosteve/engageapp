package controllers

import (
	"log"
	"net/http"
	"os"

	"github.com/engageapp/pkg/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Helper struct {
	*utils.Helper
}

// Will contain the code that handles api endpoints

// RunBroker is to run broker service
func (b *Base) RunBroker() {

	port := ":80"

	b.Router = chi.NewRouter()

	setCors(b.Router)

	// Check if the broker works
	b.Router.Use(middleware.Heartbeat("/ping-broker"))

	// Routes
	b.Router.Post("/broker", b.Broker)

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

func (b *Base) RunAuth() {
	port := ":81"
	b.Router = chi.NewRouter()

	setCors(b.Router)

	b.Router.Use((middleware.Heartbeat("/ping-auth")))

	log.Printf("Running auth service on port %s ...\n", port)
	srv := &http.Server{
		Addr:    port,
		Handler: b.Router,
	}

	// Routes
	b.Router.Post("/register", b.PostUser)

	err := srv.ListenAndServe()
	if err != nil {
		log.Printf("Error %s ", err)
		os.Exit(2)
	}
}

func setCors(r *chi.Mux) {
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"POST", "GET", "DELETE", "PUT", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

}
