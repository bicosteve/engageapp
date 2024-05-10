package entities

import "net/http"

type Config struct {
	App *http.Handler
}
