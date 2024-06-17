package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/engageapp/pkg/entities"
)

// type PostModel entities.PostModel

// func (p *PostModel) CreatePost(post *entities.PostPayload, userId int, db *sql.DB) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()

// 	query := `INSERT INTO post(message, created_at, updated_at, user_id)
// 		  VALUES(?, ?, ?, ?)`

// 	data := []interface{}{post.Message, time.Now(), time.Now(), userId}

// 	_, err := db.ExecContext(ctx, query, data...)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func CreatePost(p entities.PostValidator, userId int, db *sql.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := p.ValidatePost()
	if err != nil {
		return err
	}

	query := `INSERT INTO post(message, created_at, updated_at, user_id)
		  VALUES(?, ?, ?, ?)`

	data := []interface{}{p.GetMessage(), time.Now(), time.Now(), userId}

	_, err = db.ExecContext(ctx, query, data...)

	if err != nil {
		return err
	}

	return nil

}
