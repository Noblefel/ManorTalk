package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Noblefel/ManorTalk/backend/internal/models"
	service "github.com/Noblefel/ManorTalk/backend/internal/service/user"
	res "github.com/Noblefel/ManorTalk/backend/internal/utils/response"
	"github.com/Noblefel/ManorTalk/backend/internal/utils/validate"
	"github.com/go-chi/chi/v5"
	"github.com/gosimple/slug"
)

type UserHandlers struct {
	service service.UserService
}

func NewUserHandlers(s service.UserService) *UserHandlers {
	return &UserHandlers{
		service: s,
	}
}

func (h *UserHandlers) CheckUsername(w http.ResponseWriter, r *http.Request) {
	var payload models.CheckUsernameInput

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		res.MessageJSON(w, http.StatusBadRequest, "Error decoding json")
		return
	}

	username := slug.Make(payload.Username)

	if err := validate.Struct(payload); err != nil {
		res.JSON(w, http.StatusBadRequest, res.Response{
			Message: "Username is invalid",
			Errors:  err,
		})
		return
	}

	err := h.service.CheckUsername(username)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrDuplicateUsername):
			res.MessageJSON(w, http.StatusConflict, err.Error())
			return
		default:
			log.Println(err)
			res.MessageJSON(w, http.StatusInternalServerError, "Sorry, we had some problems checking usernames")
			return
		}
	}

	res.MessageJSON(w, http.StatusOK, "Username is available")
}

func (h *UserHandlers) Get(w http.ResponseWriter, r *http.Request) {
	user, err := h.service.Get(chi.URLParam(r, "username"))
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNoUser):
			res.MessageJSON(w, http.StatusNotFound, err.Error())
			return
		default:
			log.Println(err)
			res.MessageJSON(w, http.StatusInternalServerError, "Sorry, we had some problems retrieving the user")
			return
		}
	}

	res.JSON(w, http.StatusOK, res.Response{
		Data: user,
	})
}
