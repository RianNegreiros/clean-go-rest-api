package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/RianNegreiros/clean-go-rest-api/internal/comment"
	"github.com/google/uuid"
)

type CommentRow struct {
	ID     string
	Slug   sql.NullString
	Body   sql.NullString
	Author sql.NullString
}

func convertCommentRowToComment(cmtRow CommentRow) comment.Comment {
	return comment.Comment{
		ID:   cmtRow.ID,
		Slug: cmtRow.Slug.String,
		Body: cmtRow.Body.String,
	}
}

func (d *Database) GetComment(
	ctx context.Context,
	uuid string,
) (comment.Comment, error) {
	var cmtRow CommentRow
	row := d.Client.QueryRowContext(
		ctx,
		`SELECT id, slug, body, author
		FROM comments
		WHERE id = $1`,
		uuid,
	)
	err := row.Scan(&cmtRow.ID, &cmtRow.Slug, &cmtRow.Body, &cmtRow.Author)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("error scanning comment row: %w", err)
	}

	return convertCommentRowToComment(cmtRow), nil
}

func (d *Database) PostComment(ctx context.Context, cmt comment.Comment) (comment.Comment, error) {
	cmt.ID = uuid.New().String()
	postRow := CommentRow{
		ID:     cmt.ID,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
	}
	rows, err := d.Client.NamedQueryContext(
		ctx,
		`INSERT INTO comments 
		(id, slug, body, author)
		VALUES (:id, :slug, :body, :author)`,
		postRow,
	)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("error inserting comment: %w", err)
	}
	if err := rows.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("error closing rows: %w", err)
	}

	return cmt, nil
}
