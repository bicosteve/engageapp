package models

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/engageapp/pkg/entities"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func Register(user entities.UserValidator, p *entities.UserPayload, db *sql.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := user.ValidateUser()
	if err != nil {
		return err
	}

	hash, err := user.HashPassword()
	if err != nil {
		return err
	}

	q := `INSERT INTO user (email,password_hash,created_at,updated_at)
		  VALUES (?, ?, ?, ?)`

	data := []interface{}{p.Email, hash, time.Now(), time.Now()}

	_, err = db.ExecContext(ctx, q, data...)

	if err != nil {
		mysqlErr, ok := err.(*mysql.MySQLError)
		if mysqlErr.Number == 1062 && strings.Contains(mysqlErr.Message, "Duplicate entry") {
			return errors.New("email already in use")
		}

		_ = ok

		return err
	}

	return nil
}

func Login(u entities.UserValidator, p *entities.UserPayload, db *sql.DB) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := u.ValidateLogins()
	if err != nil {
		return "", err
	}

	var dbUser entities.User
	var user entities.User

	q := `SELECT * FROM user WHERE email = ? LIMIT 1`
	row := db.QueryRowContext(ctx, q, p.Email)
	err = row.Scan(
		&dbUser.ID,
		&dbUser.Email,
		&dbUser.Password,
		&dbUser.CreatedAt,
		&dbUser.UpdatedAt,
	)

	if err != nil {
		return "", err
	}

	requestPassword := []byte(p.Password)
	hashedPassword := []byte(dbUser.Password)

	err = bcrypt.CompareHashAndPassword(hashedPassword, requestPassword)
	if err != nil {
		return "", errors.New("password and email do not match")
	}

	token, err := user.GenerateAuthToken()
	if err != nil {
		return "", err
	}

	return token, nil
}

func ValidClaim(user entities.UserValidator, r *http.Request) (int, error) {

	userId, err := user.ValidateClaims(r)
	if err != nil {
		return 0, errors.New(err.Error())
	}
	return userId, nil
}
