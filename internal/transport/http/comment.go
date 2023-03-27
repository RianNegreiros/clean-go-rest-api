package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/RianNegreiros/clean-go-rest-api/internal/comment"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type CommentService interface {
	PostComment(context.Context, comment.Comment) (comment.Comment, error)
	GetComment(ctx context.Context, ID string) (comment.Comment, error)
	UpdateComment(ctx context.Context, ID string, newCmt comment.Comment) (comment.Comment, error)
	DeleteComment(ctx context.Context, ID string) error
}

type Response struct {
	Message string
}

type PostCommentRequest struct {
	Slug   string `json:"slug" validate:"required"`
	Author string `json:"author" validate:"required"`
	Body   string `json:"body" validate:"required"`
}

func convertPostCommentRequestToComment(req PostCommentRequest) comment.Comment {
	return comment.Comment{
		Slug:   req.Slug,
		Author: req.Author,
		Body:   req.Body,
	}
}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var cmt PostCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	validate := validator.New()
	if err := validate.Struct(cmt); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	convertComment := convertPostCommentRequestToComment(cmt)

	postCommentRequest, err := h.Service.PostComment(r.Context(), convertComment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(postCommentRequest); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cmt, err := h.Service.GetComment(r.Context(), id)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var cmt comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cmt, err := h.Service.UpdateComment(r.Context(), id, cmt)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.Service.DeleteComment(r.Context(), id); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Successfully deleted"}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
