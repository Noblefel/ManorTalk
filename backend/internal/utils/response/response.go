package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors,omitempty"`
}

func JSON(w http.ResponseWriter, code int, res Response) {
	jsonBytes, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(jsonBytes)
}

// Message is a small wrapper around res.JSON
func Message(w http.ResponseWriter, code int, msg string) {
	JSON(w, code, Response{Message: msg})
}
