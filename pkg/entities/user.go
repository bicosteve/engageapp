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
	ValidateUser(payload *UserPayload) error
	ValidateLogins(payload *UserPayload) error
	HashPassword(payload *UserPayload) (string, error)
	GenerateAuthToken(user *User) (string, error)
	ValidateClaims(r *http.Request) (int, error)
}

func (um UserModel) ValidateUser(payload *UserPayload) error {
	if payload.Email == "" {
		return errors.New("email address is required")
	}

	fmt.Println(payload.Email)

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

func (um UserModel) ValidateLogins(payload *UserPayload) error {
	if payload.Email == "" {
		return errors.New("email address is required")
	}

	if payload.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

func (um UserModel) HashPassword(payload *UserPayload) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		return "", errors.New("could not generate password hash")
	}

	return string(bytes), nil
}

func (um UserModel) GenerateAuthToken(user *User) (string, error) {
	secret := []byte(os.Getenv("JWTSECRET"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"issue":   "auth service",
		"expires": jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		"user_id": string(user.ID),
		"email":   string(user.Email),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (um UserModel) ValidateClaims(r *http.Request) (int, error) {
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
