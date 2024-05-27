package entities

import (
	"errors"
	"time"
)

type Post struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
	UpdateAt  time.Time `json:"updateAt"`
	UserId    int       `json:"userId"`
}

type PostPayload struct {
	Message string `json:"message"`
}

func ValidatePost(post *PostPayload) error {
	if post.Message == "" {
		return errors.New("message is required")
	}

	return nil
}
