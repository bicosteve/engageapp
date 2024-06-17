package entities

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
	"unicode/utf8"

	"github.com/golang-jwt/jwt/v5"
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

type CustomClaims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// type UserModel struct{}

// var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type UserValidator interface {
	ValidateUser() error
	ValidateLogins() error
	HashPassword() (string, error)
	GetEmail() string
	GetPassword() string
}

func (p *UserPayload) ValidateUser() error {
	if p.Email == "" {
		return errors.New("email address is required")
	}

	if p.Password == "" {
		return errors.New("password is required")
	}

	if p.ConfirmPassword == "" {
		return errors.New("confirm password is required")
	}

	if utf8.RuneCountInString(p.ConfirmPassword) != utf8.RuneCountInString(p.Password) {
		return errors.New("confirm password is required")

	}

	return nil

}

func (p *UserPayload) ValidateLogins() error {
	if p.Email == "" {
		return errors.New("email address is required")
	}

	if p.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

func (p *UserPayload) HashPassword() (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)

	if err != nil {
		return "", errors.New("could not generate password hash")
	}

	return string(bytes), nil
}

func (p *UserPayload) GetEmail() string {
	return p.Email
}

func (p *UserPayload) GetPassword() string {
	return p.Password
}

func CreateJWTToken(user *User) (string, error) {
	secret := []byte(os.Getenv("JWTSECRET"))
	claims := jwt.MapClaims{
		"userId": user.ID,
		"email":  user.Email,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (user *User) GetTokenString(r *http.Request) (string, error) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			return "", errors.New(err.Error())
		}
		return "", err
	}

	return cookie.Value, nil
}

func (user *User) ValidateClaim(tokenStr string) (*jwt.Token, error) {
	secret := []byte(os.Getenv("JWTSECRET"))
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

}
