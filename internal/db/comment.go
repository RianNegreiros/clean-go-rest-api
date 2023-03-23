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

func (d *Database) CreateComment(ctx context.Context, cmt comment.Comment) (comment.Comment, error) {
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

func (d *Database) DeleteComment(ctx context.Context, id string) error {
	_, err := d.Client.ExecContext(
		ctx,
		`DELETE FROM comments where id = $1`,
		id,
	)
	if err != nil {
		return fmt.Errorf("error deleting comment: %w", err)
	}

	return nil
}

func (d *Database) UpdateComment(
	ctx context.Context,
	id string,
	cmt comment.Comment,
) (comment.Comment, error) {
	cmtRow := CommentRow{
		ID:     id,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
	}

	rows, err := d.Client.NamedQueryContext(
		ctx,
		`UPDATE comments
		SET slug = :slug, body = :body, author = :author
		WHERE id = :id`,
		cmtRow,
	)

	if err != nil {
		return comment.Comment{}, fmt.Errorf("error updating comment: %w", err)
	}
	if err := rows.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("error closing rows: %w", err)
	}

	return convertCommentRowToComment(cmtRow), nil
}
