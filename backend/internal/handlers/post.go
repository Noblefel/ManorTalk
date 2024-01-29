package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/Noblefel/ManorTalk/backend/internal/models"
	service "github.com/Noblefel/ManorTalk/backend/internal/service/post"
	res "github.com/Noblefel/ManorTalk/backend/internal/utils/response"
	"github.com/Noblefel/ManorTalk/backend/internal/utils/validate"
	"github.com/go-chi/chi/v5"
)

type PostHandlers struct {
	service service.PostService
}

func NewPostHandlers(s service.PostService) *PostHandlers {
	return &PostHandlers{
		service: s,
	}
}

func (h *PostHandlers) Create(w http.ResponseWriter, r *http.Request) {
	var payload models.PostCreateInput

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		res.JSON(w, http.StatusBadRequest, res.Response{
			Message: "Error decoding json",
		})
		return
	}

	payload.Title = strings.TrimSpace(payload.Title)
	payload.Content = strings.TrimSpace(payload.Content)

	if err := validate.Struct(payload); err != nil {
		res.JSON(w, http.StatusBadRequest, res.Response{
			Message: "Some fields are invalid",
			Errors:  err,
		})
		return
	}

	post, err := h.service.Create(payload)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNoCategory):
			res.JSON(w, http.StatusNotFound, res.Response{
				Message: err.Error(),
			})
			return
		case errors.Is(err, service.ErrDuplicateTitle):
			res.JSON(w, http.StatusConflict, res.Response{
				Message: err.Error(),
			})
			return
		default:
			log.Println(err)
			res.JSON(w, http.StatusInternalServerError, res.Response{
				Message: "Sorry, we had some problems creating this post",
			})
			return
		}
	}

	res.JSON(w, http.StatusCreated, res.Response{
		Message: "Post has been created",
		Data:    post,
	})
}

func (h *PostHandlers) Get(w http.ResponseWriter, r *http.Request) {
	post, err := h.service.Get(chi.URLParam(r, "slug"))
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNoPost):
			res.JSON(w, http.StatusNotFound, res.Response{
				Message: err.Error(),
			})
			return
		default:
			log.Println(err)
			res.JSON(w, http.StatusInternalServerError, res.Response{
				Message: "Sorry, we had some problems retrieving the post",
			})
			return
		}
	}

	res.JSON(w, http.StatusOK, res.Response{
		Data: post,
	})
}

func (h *PostHandlers) GetMany(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	posts, pgMeta, err := h.service.GetMany(q)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNoCategory):
			res.JSON(w, http.StatusNotFound, res.Response{
				Message: err.Error(),
			})
			return
		default:
			log.Println(err)
			res.JSON(w, http.StatusInternalServerError, res.Response{
				Message: "Sorry, we had some problems retrieving the posts",
			})
			return
		}
	}

	res.JSON(w, http.StatusOK, res.Response{
		Data: map[string]interface{}{
			"pagination_meta": pgMeta,
			"posts":           posts,
		},
	})
}

func (h *PostHandlers) Update(w http.ResponseWriter, r *http.Request) {
	var payload models.PostUpdateInput

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		res.JSON(w, http.StatusBadRequest, res.Response{
			Message: "Error decoding json",
		})
		return
	}

	payload.Title = strings.TrimSpace(payload.Title)
	payload.Content = strings.TrimSpace(payload.Content)

	if err := validate.Struct(payload); err != nil {
		res.JSON(w, http.StatusBadRequest, res.Response{
			Message: "Some fields are invalid",
			Errors:  err,
		})
		return
	}

	err := h.service.Update(payload, chi.URLParam(r, "slug"))
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNoPost):
			res.JSON(w, http.StatusNotFound, res.Response{
				Message: err.Error(),
			})
			return
		case errors.Is(err, service.ErrDuplicateTitle):
			res.JSON(w, http.StatusConflict, res.Response{
				Message: err.Error(),
			})
			return
		default:
			log.Println(err)
			res.JSON(w, http.StatusInternalServerError, res.Response{
				Message: "Sorry, we had some problems updating the post",
			})
			return
		}
	}

	res.JSON(w, http.StatusOK, res.Response{
		Message: "Post has been updated",
	})
}

func (h *PostHandlers) Delete(w http.ResponseWriter, r *http.Request) {

	err := h.service.Delete(chi.URLParam(r, "slug"))
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNoPost):
			res.JSON(w, http.StatusNotFound, res.Response{
				Message: err.Error(),
			})
			return
		default:
			log.Println(err)
			res.JSON(w, http.StatusInternalServerError, res.Response{
				Message: "Sorry, we had some problems deleting the post",
			})
			return
		}
	}

	res.JSON(w, http.StatusOK, res.Response{
		Message: "Post has been deleted",
	})
}

func (h *PostHandlers) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.service.GetCategories()
	if err != nil {
		log.Println(err)
		res.JSON(w, http.StatusInternalServerError, res.Response{
			Message: "Sorry, we had some problems retrieving categories",
		})
		return
	}

	res.JSON(w, http.StatusOK, res.Response{
		Data: categories,
	})
}
