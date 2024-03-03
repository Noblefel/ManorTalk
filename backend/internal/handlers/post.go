package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Noblefel/ManorTalk/backend/internal/models"
	service "github.com/Noblefel/ManorTalk/backend/internal/service/post"
	res "github.com/Noblefel/ManorTalk/backend/internal/utils/response"
	"github.com/Noblefel/ManorTalk/backend/internal/utils/validate"
	"github.com/go-chi/chi/v5"
	"github.com/gosimple/slug"
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
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		res.MessageJSON(w, http.StatusBadRequest, "Error parsing form")
		return
	}

	cId, _ := strconv.Atoi(r.FormValue("category_id"))

	payload := models.PostCreateInput{
		Title:      strings.TrimSpace(r.FormValue("title")),
		Content:    strings.TrimSpace(r.FormValue("content")),
		CategoryId: cId,
		Files:      r.MultipartForm.File,
	}
	payload.Slug = slug.Make(payload.Title)

	if err := validate.Struct(payload); err != nil {
		res.JSON(w, http.StatusBadRequest, res.Response{
			Message: "Some fields are invalid",
			Errors:  err,
		})
		return
	}

	userId := r.Context().Value("user_id").(int)

	post, err := h.service.Create(payload, userId)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNoCategory):
			res.MessageJSON(w, http.StatusNotFound, err.Error())
			return
		case errors.Is(err, service.ErrImageTooLarge), errors.Is(err, service.ErrImageInvalid):
			res.MessageJSON(w, http.StatusBadRequest, err.Error())
			return
		case errors.Is(err, service.ErrDuplicateTitle):
			res.MessageJSON(w, http.StatusConflict, err.Error())
			return
		default:
			log.Println(err)
			res.MessageJSON(w, http.StatusInternalServerError, "Sorry, we had some problems creating this post")
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
			res.MessageJSON(w, http.StatusNotFound, err.Error())
			return
		default:
			log.Println(err)
			res.MessageJSON(w, http.StatusInternalServerError, "Sorry, we had some problems retrieving the post")
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
			res.MessageJSON(w, http.StatusNotFound, err.Error())
			return
		default:
			log.Println(err)
			res.MessageJSON(w, http.StatusInternalServerError, "Sorry, we had some problems retrieving the posts")
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
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		res.MessageJSON(w, http.StatusBadRequest, "Error parsing form")
		return
	}

	cId, _ := strconv.Atoi(r.FormValue("category_id"))

	payload := models.PostUpdateInput{
		Title:      strings.TrimSpace(r.FormValue("title")),
		Content:    strings.TrimSpace(r.FormValue("content")),
		CategoryId: cId,
		Files:      r.MultipartForm.File,
	}
	payload.Slug = slug.Make(payload.Title)

	if err := validate.Struct(payload); err != nil {
		res.JSON(w, http.StatusBadRequest, res.Response{
			Message: "Some fields are invalid",
			Errors:  err,
		})
		return
	}

	authId := r.Context().Value("user_id").(int)

	err := h.service.Update(payload, chi.URLParam(r, "slug"), authId)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNoPost):
			res.MessageJSON(w, http.StatusNotFound, err.Error())
			return
		case errors.Is(err, service.ErrNoCategory):
			res.MessageJSON(w, http.StatusNotFound, err.Error())
			return
		case errors.Is(err, service.ErrUnauthorized):
			res.MessageJSON(w, http.StatusUnauthorized, err.Error())
			return
		case errors.Is(err, service.ErrImageTooLarge), errors.Is(err, service.ErrImageInvalid):
			res.MessageJSON(w, http.StatusBadRequest, err.Error())
			return
		case errors.Is(err, service.ErrDuplicateTitle):
			res.MessageJSON(w, http.StatusConflict, err.Error())
			return
		default:
			log.Println(err)
			res.MessageJSON(w, http.StatusInternalServerError, "Sorry, we had some problems updating the post")
			return
		}
	}

	res.JSON(w, http.StatusOK, res.Response{
		Message: "Post has been updated",
		Data:    payload.Slug,
	})
}

func (h *PostHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	authId := r.Context().Value("user_id").(int)

	err := h.service.Delete(chi.URLParam(r, "slug"), authId)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNoPost):
			res.MessageJSON(w, http.StatusNotFound, err.Error())
			return
		case errors.Is(err, service.ErrUnauthorized):
			res.MessageJSON(w, http.StatusUnauthorized, err.Error())
			return
		default:
			log.Println(err)
			res.MessageJSON(w, http.StatusInternalServerError, "Sorry, we had some problems deleting the post")
			return
		}
	}

	res.MessageJSON(w, http.StatusOK, "Post has been deleted")
}

func (h *PostHandlers) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.service.GetCategories()
	if err != nil {
		log.Println(err)
		res.MessageJSON(w, http.StatusInternalServerError, "Sorry, we had some problems retrieving categories")
		return
	}

	res.JSON(w, http.StatusOK, res.Response{
		Data: categories,
	})
}
