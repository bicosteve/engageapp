package models

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/engageapp/pkg/entities"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct{}

func (um *UserModel) RegisterUser(user *entities.UserPayload, db *sql.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	hash, err := entities.HashPassword(user)

	if err != nil {
		return err
	}

	q := `INSERT INTO user (email,password_hash,created_at,updated_at)
		  VALUES (?, ?, ?, ?)`

	data := []interface{}{user.Email, hash, time.Now(), time.Now()}

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

func (um *UserModel) LoginUser(user *entities.UserPayload, db *sql.DB) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var dbUser entities.User

	q := `SELECT * FROM user WHERE email = ? LIMIT 1`
	row := db.QueryRowContext(ctx, q, user.Email)
	err := row.Scan(
		&dbUser.ID,
		&dbUser.Email,
		&dbUser.Password,
		&dbUser.CreatedAt,
		&dbUser.UpdatedAt,
	)

	if err != nil {
		return "", err
	}

	requestPassword := []byte(user.Password)
	hashedPassword := []byte(dbUser.Password)

	err = bcrypt.CompareHashAndPassword(hashedPassword, requestPassword)
	if err != nil {
		return "", errors.New("password and email do not match")
	}

	token, err := entities.GenerateAuthToken(&dbUser)
	if err != nil {
		return "", err
	}

	return token, nil
}
