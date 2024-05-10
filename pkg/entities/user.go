package entities

import (
	"errors"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserPayload struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type Claims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func Validate(payload *UserPayload) error {
	if payload.Email == "" {
		return errors.New("email address is required")
	}

	if payload.Password == "" {
		return errors.New("password is required")
	}

	if payload.ConfirmPassword == "" {
		return errors.New("confirm password is required")
	}

	if payload.ConfirmPassword != payload.Password {
		return errors.New("confirm password must match password")
	}

	return nil
}
