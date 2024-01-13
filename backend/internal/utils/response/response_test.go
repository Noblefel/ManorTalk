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
	{
		name: "response-ok",
		payload: Response{
			Message: "Success",
		},
		inputStatusCode:    http.StatusOK,
		expectedStatusCode: http.StatusOK,
	},
	{
		name: "response-ok-2",
		payload: Response{
			Message: "Some fields are invalid",
		},
		inputStatusCode:    http.StatusBadRequest,
		expectedStatusCode: http.StatusBadRequest,
	},
	{
		name: "response-error-marshalling-json",
		payload: Response{
			Data: make(chan int),
		},
		inputStatusCode:    http.StatusOK,
		expectedStatusCode: http.StatusInternalServerError,
	},
}

func TestJSON(t *testing.T) {
	for _, tt := range jsonTests {
		r := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()

		JSON(w, r, tt.inputStatusCode, tt.payload)

		if w.Code != tt.expectedStatusCode {
			t.Errorf("%s returned response code of %d, wanted %d", tt.name, w.Code, tt.expectedStatusCode)
		}
	}
}
