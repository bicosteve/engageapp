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

type PostModel struct{}

type PostPayload struct {
	Message string `json:"message"`
}

type PostValidator interface {
	ValidatePost() error
}

// func ValidatePost(post *PostPayload) error {
// 	if post.Message == "" {
// 		return errors.New("message is required")
// 	}

// 	return nil
// }

func (p *PostPayload) ValidatePost() error {
	if p.Message == "" {
		return errors.New("message is required")
	}

	return nil
}
