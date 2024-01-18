package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/Noblefel/ManorTalk/backend/internal/config"
	"github.com/Noblefel/ManorTalk/backend/internal/database"
	"github.com/Noblefel/ManorTalk/backend/internal/models"
	"github.com/Noblefel/ManorTalk/backend/internal/repository"
	"github.com/Noblefel/ManorTalk/backend/internal/repository/postgres"
	"github.com/Noblefel/ManorTalk/backend/internal/repository/redis"
	res "github.com/Noblefel/ManorTalk/backend/internal/utils/response"
	"github.com/Noblefel/ManorTalk/backend/internal/utils/validate"
	"github.com/go-chi/chi/v5"

	"github.com/gosimple/slug"
)

type PostHandlers struct {
	c         *config.AppConfig
	redisRepo repository.RedisRepo
	postRepo  repository.PostRepo
}

func NewPostHandlers(c *config.AppConfig, db *database.DB) *PostHandlers {
	return &PostHandlers{
		c:         c,
		redisRepo: redis.NewRepo(db),
		postRepo:  postgres.NewPostRepo(db),
	}
}

func NewTestPostHandlers(c *config.AppConfig) *PostHandlers {
	return &PostHandlers{
		c:         c,
		redisRepo: redis.NewTestRepo(),
		postRepo:  postgres.NewTestPostRepo(),
	}
}

func (h *PostHandlers) Create(w http.ResponseWriter, r *http.Request) {
	var payload models.PostCreateInput

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		res.JSON(w, r, http.StatusBadRequest, res.Response{
			Message: "Error decoding json",
		})
		return
	}

	payload.Title = strings.TrimSpace(payload.Title)
	payload.Content = strings.TrimSpace(payload.Content)

	if err := validate.Struct(payload); err != nil {
		res.JSON(w, r, http.StatusBadRequest, res.Response{
			Message: "Some fields are invalid",
			Errors:  err,
		})
		return
	}

	post := models.Post{
		UserId:     1,
		Title:      payload.Title,
		Slug:       slug.Make(payload.Title),
		Excerpt:    payload.Excerpt,
		Content:    payload.Content,
		CategoryId: payload.CategoryId,
	}

	post, err := h.postRepo.CreatePost(post)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			res.JSON(w, r, http.StatusConflict, res.Response{
				Message: "Title has already been used",
			})
			return
		}

		res.JSON(w, r, http.StatusInternalServerError, res.Response{
			Message: "Unexpected error when creating the post",
		})
		return
	}

	res.JSON(w, r, http.StatusCreated, res.Response{
		Message: "Post has been created",
		Data:    post,
	})
}

func (h *PostHandlers) Get(w http.ResponseWriter, r *http.Request) {
	post, err := h.postRepo.GetPostBySlug(chi.URLParam(r, "slug"))
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			res.JSON(w, r, http.StatusNotFound, res.Response{
				Message: "Post not found",
			})
			return
		}

		res.JSON(w, r, http.StatusInternalServerError, res.Response{
			Message: "Unexpected error when getting post",
		})
		return
	}

	res.JSON(w, r, http.StatusOK, res.Response{
		Data: post,
	})
}

func (h *PostHandlers) Update(w http.ResponseWriter, r *http.Request) {
	var payload models.PostUpdateInput

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		res.JSON(w, r, http.StatusBadRequest, res.Response{
			Message: "Error decoding json",
		})
		return
	}

	payload.Title = strings.TrimSpace(payload.Title)
	payload.Content = strings.TrimSpace(payload.Content)

	if err := validate.Struct(payload); err != nil {
		res.JSON(w, r, http.StatusBadRequest, res.Response{
			Message: "Some fields are invalid",
			Errors:  err,
		})
		return
	}

	post, err := h.postRepo.GetPostBySlug(chi.URLParam(r, "slug"))
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			res.JSON(w, r, http.StatusNotFound, res.Response{
				Message: "Post not found",
			})
			return
		}

		res.JSON(w, r, http.StatusInternalServerError, res.Response{
			Message: "Unexpected error when getting the post",
		})
		return
	}

	post = models.Post{
		Id:         post.Id,
		Title:      payload.Title,
		Slug:       slug.Make(payload.Title),
		Excerpt:    payload.Excerpt,
		Content:    payload.Content,
		CategoryId: payload.CategoryId,
	}

	err = h.postRepo.UpdatePost(post)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			res.JSON(w, r, http.StatusConflict, res.Response{
				Message: "Title has already been used",
			})
			return
		}

		res.JSON(w, r, http.StatusInternalServerError, res.Response{
			Message: "Unexpected error when updating the post",
		})
		return
	}

	res.JSON(w, r, http.StatusOK, res.Response{
		Message: "Post has been updated",
	})
}

func (h *PostHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	post, err := h.postRepo.GetPostBySlug(chi.URLParam(r, "slug"))
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			res.JSON(w, r, http.StatusNotFound, res.Response{
				Message: "Post not found",
			})
			return
		}

		res.JSON(w, r, http.StatusInternalServerError, res.Response{
			Message: "Unexpected error when getting the post",
		})
		return
	}

	err = h.postRepo.DeletePost(post.Id)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Response{
			Message: "Unexpected error when deleting the post",
		})
		return
	}

	res.JSON(w, r, http.StatusOK, res.Response{
		Message: "Post has been deleted",
	})
}
