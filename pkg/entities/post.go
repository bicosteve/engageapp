package entities

import "time"

type Post struct {
	id        int       `json:"id"`
	message   string    `json:"message"`
	createdAt time.Time `json:"createdAt"`
	updateAt  time.Time `json:"updateAt"`
	userId    int       `json:"userId"`
}

type PostPayload struct {
	message string `json:"message"`
	userId  int    `json:"userId"`
}
