package entities

import (
	"database/sql"
	"errors"
	"time"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"
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
	ConfirmPassword string `json:"confirm_password,omitempty"`
}

type UserModel struct {
	DB *sql.DB
}

type Claims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

var errMap = make(map[string]string)

// var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

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

	if utf8.RuneCountInString(payload.ConfirmPassword) != utf8.RuneCountInString(payload.Password) {
		return errors.New("confirm password is required")

	}

	return nil

}

func ValidateLogin(payload *UserPayload) error {
	if payload.Email == "" {
		return errors.New("email address is required")
	}

	if payload.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

func HashPassword(payload *UserPayload) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		return "", errors.New("could not generate password hash")
	}

	return string(bytes), nil
}
