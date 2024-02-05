package handlers

import (
	"net/http"

	res "github.com/Noblefel/ManorTalk/backend/internal/utils/response"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	res.MessageJSON(w, http.StatusNotFound, "Not Found")
}

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	res.MessageJSON(w, http.StatusMethodNotAllowed, "Method Not Allowed")
}
