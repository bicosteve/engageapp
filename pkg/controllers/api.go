package controllers

import (
	"log"
	"net/http"
	"os"

	"github.com/engageapp/pkg/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Will contain the code that handles api endpoints

// RunBroker is to run broker service
func (b *Base) RunBroker() {

	port := ":80"

	b.Router = chi.NewRouter()

	// set cors for chi
	utils.SetCors(b.Router)

	// Check if the broker works
	b.Router.Use(middleware.Heartbeat("/ping-broker"))

	b.Router.Post("/broker", b.Broker)

	log.Printf("Running broker service on port %s ... \n", port)
	srv := &http.Server{
		Addr:    port,
		Handler: b.Router,
	}

	// err := http.ListenAndServe(port, b.Router)

	err := srv.ListenAndServe()
	if err != nil {
		log.Printf("Error running broker %s", err)
		os.Exit(1)
	}

}

// Run Auth Service
func (b *Base) RunAuth() {
	port := ":81"
	b.Router = chi.NewRouter()

	utils.SetCors(b.Router)

	b.Router.Use((middleware.Heartbeat("/ping-auth")))

	log.Printf("Running auth service on port %s ...\n", port)
	srv := &http.Server{
		Addr:    port,
		Handler: b.Router,
	}

	b.Router.Post("/register", b.PostUser)
	b.Router.Post("/login", b.Login)

	err := srv.ListenAndServe()
	if err != nil {
		log.Printf("Error running auth service %s ", err)
		os.Exit(2)
	}

}

// Run Post Service
func (b *Base) RunPost() {
	port := ":82"

	b.Router = chi.NewRouter()
	utils.SetCors(b.Router)

	b.Router.Use(middleware.Heartbeat("/ping-post"))

	log.Printf("Running post service on port %s ...\n", port)
	srv := &http.Server{
		Addr:    port,
		Handler: b.Router,
	}

	b.Router.Post("/post", b.CreatePost)
	b.Router.Get("/posts", b.GetPosts)

	err := srv.ListenAndServe()
	if err != nil {
		log.Printf("Error on post service %s ", err)
		os.Exit(2)
	}

}

// Run Comment Service
func (b *Base) RunComment() {
	port := ":83"

	b.Router = chi.NewRouter()
	utils.SetCors(b.Router)

	b.Router.Use(middleware.Heartbeat("/ping-comment"))

	log.Printf("Running comment service on port %s ...\n", port)
	srv := &http.Server{
		Addr:    port,
		Handler: b.Router,
	}

	b.Router.Post("/comment/{commentid}", b.CreateComment)
	b.Router.Get("/comments", b.GetComments)

	err := srv.ListenAndServe()
	if err != nil {
		log.Printf("Error on post service %s ", err)
		os.Exit(2)
	}

}
