package response

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var jsonTests = []struct {
	name               string
	payload            Response
	inputStatusCode    int
	expectedStatusCode int
}{
	{"success", Response{}, http.StatusOK, http.StatusOK},
	{"success bad request", Response{}, http.StatusBadRequest, http.StatusBadRequest},
	{"error marshalling json", Response{Data: make(chan int)}, http.StatusOK, http.StatusInternalServerError},
}

func TestJSON(t *testing.T) {
	for _, tt := range jsonTests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			JSON(w, tt.inputStatusCode, tt.payload)

			if w.Code != tt.expectedStatusCode {
				t.Errorf("want %d, got %d", w.Code, tt.expectedStatusCode)
			}
		})
	}
}

func TestMessage(t *testing.T) {
	w := httptest.NewRecorder()

	Message(w, 200, "message")

	if w.Code != 200 {
		t.Errorf("want 200, got %d", w.Code)
	}
}
