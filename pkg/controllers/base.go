package controllers

import (
	"github.com/go-chi/chi/v5"
)

// Will contain the code that initializes app dependancies
type Base struct {
	Router *chi.Mux
}
