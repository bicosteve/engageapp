package entities

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"database/sql"
	"errors"
	"net/http"
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

type UserModel struct {
	DB *sql.DB
}

type Claims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

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

func ValidateLogins(payload *UserPayload) error {
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

func GenerateAuthToken(user *User) (string, error) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return "", err
	}
	claims := &Claims{
		ID:    user.ID,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "authservice",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// func GetToken(r *http.Request) (string, error) {
// 	cookie, err := r.Cookie("token")
// 	if err != nil {
// 		if err == http.ErrNoCookie {
// 			return "", errors.New("there is no cookie set")
// 		}

// 		return "", err
// 	}

// 	tokenStr := cookie.Value
// 	return tokenStr, nil

// }

func ValidateClaims(claims *Claims, r *http.Request) (*Claims, error) {
	key, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	derBuf, _ := x509.MarshalECPrivateKey(key)
	// privKey, _ := x509.ParseECPrivateKey(derBuf)

	c, err := r.Cookie("token")
	if err != nil {
		return &Claims{}, err
	}

	tkn, err := jwt.ParseWithClaims(c.Value, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(derBuf), nil
	})

	if err != nil {
		return &Claims{}, err
	}

	if !tkn.Valid {
		return &Claims{}, nil
	}

	return claims, nil
}
