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

func Register(user entities.UserValidator, db *sql.DB) error {
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

	data := []interface{}{user.GetEmail(), hash, time.Now(), time.Now()}

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

func LoginToken(u entities.UserValidator, db *sql.DB) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := u.ValidateLogins()
	if err != nil {
		return "", err
	}

	var dbUser entities.User

	q := `SELECT * FROM user WHERE email = ? LIMIT 1`
	row := db.QueryRowContext(ctx, q, u.GetEmail())
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

	requestPassword := []byte(u.GetPassword())
	hashedPassword := []byte(dbUser.Password)

	err = bcrypt.CompareHashAndPassword(hashedPassword, requestPassword)
	if err != nil {
		return "", errors.New("password and email do not match")
	}

	token, err := entities.CreateJWTToken(&dbUser)
	if err != nil {
		return "", errors.New(err.Error())
	}

	return token, nil
}

func GetByEmail(u entities.UserValidator, db *sql.DB) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var user entities.User

	err := u.ValidateLogins()
	if err != nil {
		return false, err
	}

	q := `SELECT email FROM user WHERE email = ? LIMIT 1`
	row := db.QueryRowContext(ctx, q, u.GetEmail())
	err = row.Scan(&user.Email)
	if err != nil {
		return false, err
	}
	return true, nil
}
