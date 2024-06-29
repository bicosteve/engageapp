package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/engageapp/pkg/entities"
)

func AddComment(c entities.CommentValidator, userId, postId int, db *sql.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := c.ValidateComment()
	if err != nil {
		return err
	}

	q := `INSERT INTO comment (comment, created_at, updated_at, user_id, post_id)
			VALUES (?, ?, ?, ?, ?)`

	data := []interface{}{c.GetMessage(), time.Now(), time.Now(), userId, postId}
	_, err = db.ExecContext(ctx, q, data...)
	if err != nil {
		return err
	}

	return nil
}

func GetComment(postId int, db *sql.DB) ([]*entities.Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var comments []*entities.Comment
	q := `SELECT * FROM comment WHERE post_id = ? LIMIT 100`

	rows, err := db.QueryContext(ctx, q, postId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var comment entities.Comment
		err := rows.Scan(
			&comment.ID,
			&comment.Message,
			&comment.CreatedAt,
			&comment.UpdateAt,
			&comment.UserId,
			&comment.PostId,
		)

		if err != nil {
			return nil, err
		}

		comments = append(comments, &comment)

	}

	return comments, nil
}
