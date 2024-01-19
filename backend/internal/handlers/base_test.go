package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBaseHandlers(t *testing.T) {
	var tests = []struct {
		name       string
		url        string
		method     string
		handler    http.HandlerFunc
		statusCode int
	}{
		{
			name:       "not-found",
			url:        "/xmo02v3o2cm3ro",
			method:     "GET",
			handler:    NotFound,
			statusCode: http.StatusNotFound,
		},
		{
			name:       "method-not-allowed",
			url:        "/users",
			method:     "asjcaosjdcoa",
			handler:    MethodNotAllowed,
			statusCode: http.StatusMethodNotAllowed,
		},
	}

	for _, tt := range tests {
		r, _ := http.NewRequest(tt.method, tt.url, nil)
		w := httptest.NewRecorder()

		tt.handler.ServeHTTP(w, r)

		if w.Code != tt.statusCode {
			t.Errorf("%s returned response code of %d, wanted %d", tt.name, w.Code, tt.statusCode)
		}
	}
}
