package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/engageapp/pkg/entities"
)

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

func GetPosts(db *sql.DB) ([]*entities.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_ = ctx

	var posts []*entities.Post

	q := `SELECT * FROM post LIMIT 100`
	rows, err := db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var post entities.Post
		err := rows.Scan(
			&post.ID,
			&post.Message,
			&post.CreatedAt,
			&post.UpdateAt,
			&post.UserId,
		)

		if err != nil {
			return nil, err
		}

		posts = append(posts, &post)
	}

	return posts, nil
}
