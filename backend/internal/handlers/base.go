package handlers

import (
	"net/http"

	res "github.com/Noblefel/ManorTalk/backend/internal/utils/response"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	res.JSON(w, r, http.StatusNotFound, res.Response{
		Message: "Not Found",
	})
}

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	res.JSON(w, r, http.StatusMethodNotAllowed, res.Response{
		Message: "Method Not Allowed",
	})
}
