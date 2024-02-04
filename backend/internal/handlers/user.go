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
		res.JSON(w, http.StatusBadRequest, res.Response{
			Message: "Error decoding json",
		})
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
			res.JSON(w, http.StatusConflict, res.Response{
				Message: err.Error(),
			})
			return
		default:
			log.Println(err)
			res.JSON(w, http.StatusInternalServerError, res.Response{
				Message: "Sorry, we had some problems checking usernames",
			})
			return
		}
	}

	res.JSON(w, http.StatusOK, res.Response{Message: "Username is available"})
}
