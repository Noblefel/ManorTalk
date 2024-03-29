package handlers

import (
	"bytes"
	"context"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/Noblefel/ManorTalk/backend/internal/config"
	"github.com/Noblefel/ManorTalk/backend/internal/database"
	"github.com/Noblefel/ManorTalk/backend/internal/repository/postgres"
	"github.com/Noblefel/ManorTalk/backend/internal/repository/redis"
	service "github.com/Noblefel/ManorTalk/backend/internal/service/post"
	"github.com/gosimple/slug"
)

func TestNewPostHandlers(t *testing.T) {
	var db *database.DB
	var c *config.AppConfig
	cr := redis.NewRepo(db)
	pr := postgres.NewPostRepo(db)
	s := service.NewPostService(c, cr, pr)
	post := NewPostHandlers(s)

	typeString := reflect.TypeOf(post).String()

	if typeString != "*handlers.PostHandlers" {
		t.Error("NewPostHandlers() did not get the correct type, wanted *handlers.PostHandlers")
	}
}

var longText = `A long text: Lorem ipsum dolor sit amet, consectetur adipiscing
elit. Nulla posuere neque id magna pretium rutrum. Sed ornare nunc arcu.
Cras pharetra, nibh ac ultricies blandit, purus sapien mattis turpis, et
congue felis ligula sit amet mi`

func TestPost_Create(t *testing.T) {
	var tests = []struct {
		name         string
		noForm       bool
		payloadTitle string
		statusCode   int
	}{
		{"success", false, "a sample title 123", http.StatusCreated},
		{"error parsing form", true, "", http.StatusBadRequest},
		{"error validation", false, "", http.StatusBadRequest},
		{"no category", false, service.ErrNoCategory.Error(), http.StatusNotFound},
		{"error image invalid", false, service.ErrImageInvalid.Error(), http.StatusBadRequest},
		{"error image too large", false, service.ErrImageTooLarge.Error(), http.StatusBadRequest},
		{"duplicate title", false, service.ErrDuplicateTitle.Error(), http.StatusConflict},
		{"unexpected error", false, "unexpected error", http.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var b bytes.Buffer
			fw := multipart.NewWriter(&b)
			fw.WriteField("title", tt.payloadTitle)
			fw.WriteField("content", longText)
			fw.WriteField("category_id", "1")
			fw.CreateFormFile("image", "x")
			fw.Close()

			var r *http.Request
			if tt.noForm {
				r = httptest.NewRequest("POST", "/posts", nil)
			} else {
				r = httptest.NewRequest("POST", "/posts", &b)
			}

			ctx := context.WithValue(r.Context(), "user_id", 1)
			r = r.WithContext(ctx)
			r.Header.Set("Content-Type", fw.FormDataContentType())
			w := httptest.NewRecorder()
			handler := http.HandlerFunc(h.post.Create)
			handler.ServeHTTP(w, r)

			if w.Code != tt.statusCode {
				t.Errorf("want %d, got %d", tt.statusCode, w.Code)
			}
		})
	}
}

func TestPost_Get(t *testing.T) {
	var tests = []struct {
		name       string
		slugRoute  string
		statusCode int
	}{
		{"success", "post-title", http.StatusOK},
		{"no post", service.ErrNoPost.Error(), http.StatusNotFound},
		{"unexpected error", "unexpected error", http.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest("GET", "/posts/{slug}", nil)
			ctx := getCtxWithParam(r, params{"slug": tt.slugRoute})
			r = r.WithContext(ctx)
			w := httptest.NewRecorder()
			handler := http.HandlerFunc(h.post.Get)
			handler.ServeHTTP(w, r)

			if w.Code != tt.statusCode {
				t.Errorf("want %d, got %d", tt.statusCode, w.Code)
			}
		})
	}
}

func TestPost_GetMany(t *testing.T) {
	var tests = []struct {
		name       string
		query      string
		statusCode int
	}{
		{"success", "page=1&limit=10", http.StatusOK},
		{"no category", slug.Make(service.ErrNoCategory.Error()) + "=1", http.StatusNotFound},
		{"unexpected error", slug.Make("unexpected error") + "=1", http.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest("GET", "/posts?"+tt.query, nil)
			w := httptest.NewRecorder()
			handler := http.HandlerFunc(h.post.GetMany)
			handler.ServeHTTP(w, r)

			if w.Code != tt.statusCode {
				t.Errorf("want %d, got %d", tt.statusCode, w.Code)
			}
		})
	}
}

func TestPost_Update(t *testing.T) {
	var tests = []struct {
		name           string
		slugRoute      string
		noForm         bool
		failValidation bool
		statusCode     int
	}{
		{"success", "slug", false, false, http.StatusOK},
		{"error parsing form", "", true, false, http.StatusBadRequest},
		{"error validation", "", false, true, http.StatusBadRequest},
		{"no post", service.ErrNoPost.Error(), false, false, http.StatusNotFound},
		{"unauthorized", service.ErrUnauthorized.Error(), false, false, http.StatusUnauthorized},
		{"no category", service.ErrNoCategory.Error(), false, false, http.StatusNotFound},
		{"error image invalid", service.ErrImageInvalid.Error(), false, false, http.StatusBadRequest},
		{"error image too large", service.ErrImageTooLarge.Error(), false, false, http.StatusBadRequest},
		{"duplicate title", service.ErrDuplicateTitle.Error(), false, false, http.StatusConflict},
		{"unexpected error", "unexpected error", false, false, http.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var b bytes.Buffer
			fw := multipart.NewWriter(&b)
			fw.WriteField("content", longText)
			fw.WriteField("category_id", "1")
			if tt.failValidation {
				fw.WriteField("title", "x")
			} else {
				fw.WriteField("title", "sample title")
			}
			fw.CreateFormFile("image", "x")
			fw.Close()

			var r *http.Request
			if tt.noForm {
				r = httptest.NewRequest("PATCH", "/posts/{slug}", nil)
			} else {
				r = httptest.NewRequest("PATCH", "/posts/{slug}", &b)
			}

			ctx := getCtxWithParam(r, params{"slug": tt.slugRoute})
			ctx = context.WithValue(ctx, "user_id", 1)
			r = r.WithContext(ctx)
			r.Header.Set("Content-Type", fw.FormDataContentType())
			w := httptest.NewRecorder()
			handler := http.HandlerFunc(h.post.Update)
			handler.ServeHTTP(w, r)

			if w.Code != tt.statusCode {
				t.Errorf("want %d, got %d", tt.statusCode, w.Code)
			}
		})
	}
}

func TestPost_Delete(t *testing.T) {
	var tests = []struct {
		name       string
		slugRoute  string
		statusCode int
	}{
		{"success", "post-title", http.StatusOK},
		{"no post", service.ErrNoPost.Error(), http.StatusNotFound},
		{"unauthorized", service.ErrUnauthorized.Error(), http.StatusUnauthorized},
		{"unexpected error", "unexpected error", http.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest("DELETE", "/posts/{slug}", nil)
			ctx := getCtxWithParam(r, params{"slug": tt.slugRoute})
			ctx = context.WithValue(ctx, "user_id", 1)
			r = r.WithContext(ctx)
			w := httptest.NewRecorder()
			handler := http.HandlerFunc(h.post.Delete)
			handler.ServeHTTP(w, r)

			if w.Code != tt.statusCode {
				t.Errorf("want %d, got %d", tt.statusCode, w.Code)
			}
		})
	}
}

func TestPost_GetCategories(t *testing.T) {
	var tests = []struct {
		name       string
		statusCode int
	}{
		{"success", http.StatusOK},
	}

	for _, tt := range tests {
		r := httptest.NewRequest("GET", "/posts/categories", nil)
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(h.post.GetCategories)
		handler.ServeHTTP(w, r)

		if w.Code != tt.statusCode {
			t.Errorf("want %d, got %d", tt.statusCode, w.Code)
		}
	}
}
