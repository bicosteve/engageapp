package entities

import (
	"errors"
	"time"
)

type Comment struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
	UserId    int       `json:"user_id"`
	PostId    int       `json:"post_id"`
}

type CommentPayload struct {
	Message string `json:"message"`
}

type CommentValidator interface {
	ValidateComment() error
	GetMessage() string
}

func (c *CommentPayload) ValidateComment() error {
	if c.Message == "" {
		return errors.New("message cannot be empty")
	}
	return nil
}

func (c *CommentPayload) GetMessage() string {
	return c.Message
}
