package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/engageapp/pkg/entities"
	"github.com/engageapp/pkg/utils"
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
		utils.Log("ERROR", "usermdl", err.Error())
		return err
	}

	return nil

}
