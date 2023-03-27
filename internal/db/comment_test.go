//go:build integration
// +build integration

package db

import (
	"context"
	"testing"

	"github.com/RianNegreiros/clean-go-rest-api/internal/comment"

	"github.com/stretchr/testify/assert"
)

func TestCommentDatabase(t *testing.T) {
	t.Run("Create a comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug:   "test",
			Author: "test",
			Body:   "test",
		})
		assert.NoError(t, err)

		newCmt, err := db.GetComment(context.Background(), cmt.ID)
		assert.NoError(t, err)
		assert.Equal(t, "test", newCmt.Slug)
	})
}
