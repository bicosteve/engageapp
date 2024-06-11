package entities

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
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

type Claims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type UserModel struct{}

// var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type UserValidator interface {
	ValidateUser() error
	ValidateLogins() error
	HashPassword() (string, error)
	ValidateClaims(r *http.Request) (int, error)
}

func (p *UserPayload) ValidateUser() error {
	if p.Email == "" {
		return errors.New("email address is required")
	}

	fmt.Println(p.Email)

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

func (p *User) GenerateAuthToken() (string, error) {
	secret := []byte(os.Getenv("JWTSECRET"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"issue":   "auth service",
		"expires": jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		"user_id": string(p.ID),
		"email":   string(p.Email),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (p *UserPayload) ValidateClaims(r *http.Request) (int, error) {
	secret := []byte(os.Getenv("JWTSECRET"))
	myCookie, err := r.Cookie("token")
	if err != nil {
		return 0, errors.New(err.Error())
	}

	token, err := jwt.Parse(myCookie.Value, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secret, nil
	})

	if err != nil {
		return 0, errors.New(err.Error())
	}

	if token == nil {
		return 0, errors.New("there is no token in the header")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("token parse error")
	}

	mapData, ok := claims["data"].(map[string]interface{})
	if !ok {
		return 0, errors.New("assert claims[\"data\"] as map[string]interface{} failed")
	}

	exp := claims["expires"].(float64)
	if int64(exp) < time.Now().Unix() {
		return 0, errors.New("token expired")
	}

	userStr, ok := mapData["user_id"].(string)
	if !ok {
		return 0, errors.New("assert mapData[\"user_id\"] as string failed")
	}

	userID, _ := strconv.Atoi(userStr)

	return userID, nil

}
