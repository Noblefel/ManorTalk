package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/Noblefel/ManorTalk/backend/internal/config"
	"github.com/Noblefel/ManorTalk/backend/internal/database"
	"github.com/Noblefel/ManorTalk/backend/internal/models"
	"github.com/Noblefel/ManorTalk/backend/internal/repository"
)

func TestNewPostHandlers(t *testing.T) {
	var db *database.DB
	var c *config.AppConfig
	auth := NewPostHandlers(c, db)

	typeString := reflect.TypeOf(auth).String()

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
		name       string
		payload    *models.PostCreateInput
		statusCode int
	}{
		{
			name: "postCreate-ok",
			payload: &models.PostCreateInput{
				Title:      "The new post title",
				Content:    longText,
				CategoryId: 1,
			},
			statusCode: http.StatusCreated,
		},
		{
			name:       "postCreate-error-decode-json",
			payload:    nil,
			statusCode: http.StatusBadRequest,
		},
		{
			name: "postCreate-error-validation",
			payload: &models.PostCreateInput{
				Title:   "",
				Content: "",
			},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "postCreate-error-category-not-found",
			payload: &models.PostCreateInput{
				Title:      "The new post title",
				Content:    longText,
				CategoryId: repository.ErrNotFoundKeyInt,
			},
			statusCode: http.StatusNotFound,
		},
		{
			name: "postCreate-error-getting-category",
			payload: &models.PostCreateInput{
				Title:      "The new post title",
				Content:    longText,
				CategoryId: repository.ErrUnexpectedKeyInt,
			},
			statusCode: http.StatusInternalServerError,
		},
		{
			name: "postCreate-error-duplicate-title-or-slug",
			payload: &models.PostCreateInput{
				Title:      repository.ErrDuplicateKeyString,
				Content:    longText,
				CategoryId: 1,
			},
			statusCode: http.StatusConflict,
		},
		{
			name: "postCreate-error-creating-post",
			payload: &models.PostCreateInput{
				Title:      repository.ErrUnexpectedKeyString,
				Content:    longText,
				CategoryId: 1,
			},
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		var r *http.Request
		if tt.payload == nil {
			r = httptest.NewRequest("POST", "/posts", nil)
		} else {
			jsonBytes, _ := json.Marshal(tt.payload)
			r = httptest.NewRequest("POST", "/posts", bytes.NewBuffer(jsonBytes))
		}

		w := httptest.NewRecorder()
		handler := http.HandlerFunc(h.post.Create)
		handler.ServeHTTP(w, r)

		if w.Code != tt.statusCode {
			t.Errorf("%s returned response code of %d, wanted %d", tt.name, w.Code, tt.statusCode)
		}
	}
}

func TestPost_Get(t *testing.T) {
	var tests = []struct {
		name       string
		slugRoute  string
		statusCode int
	}{
		{
			name:       "postGet-ok",
			slugRoute:  "post-title",
			statusCode: http.StatusOK,
		},
		{
			name:       "postGet-error-post-not-found",
			slugRoute:  repository.ErrNotFoundKeyString,
			statusCode: http.StatusNotFound,
		},
		{
			name:       "postGet-error-getting-post",
			slugRoute:  repository.ErrUnexpectedKeyString,
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {

		r := httptest.NewRequest("GET", "/posts/{slug}", nil)
		ctx := getCtxWithParam(r, params{"slug": tt.slugRoute})
		r = r.WithContext(ctx)
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(h.post.Get)
		handler.ServeHTTP(w, r)

		if w.Code != tt.statusCode {
			t.Errorf("%s returned response code of %d, wanted %d", tt.name, w.Code, tt.statusCode)
		}
	}
}

func TestPost_GetByCategory(t *testing.T) {
	var tests = []struct {
		name       string
		category   string
		statusCode int
	}{
		{
			name:       "postGetByCategory-ok",
			category:   "other",
			statusCode: http.StatusOK,
		},
		{
			name:       "postGetByCategory-ok-even-with-invalid-category",
			category:   "asdcoashdcisaohdnashdonahscdaosha",
			statusCode: http.StatusOK,
		},
		{
			name:       "postGetByCategory-error-getting-posts",
			category:   repository.ErrUnexpectedKeyString,
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		r := httptest.NewRequest("GET", "/posts/c/{category}", nil)
		ctx := getCtxWithParam(r, params{"category": tt.category})
		r = r.WithContext(ctx)
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(h.post.GetByCategory)
		handler.ServeHTTP(w, r)

		if w.Code != tt.statusCode {
			t.Errorf("%s returned response code of %d, wanted %d", tt.name, w.Code, tt.statusCode)
		}
	}
}

func TestPost_Update(t *testing.T) {
	var tests = []struct {
		name       string
		slugRoute  string
		payload    *models.PostUpdateInput
		statusCode int
	}{
		{
			name:      "postUpdate-ok",
			slugRoute: "post-title",
			payload: &models.PostUpdateInput{
				Title:      "The updated post title",
				Content:    longText,
				CategoryId: 1,
			},
			statusCode: http.StatusOK,
		},
		{
			name:       "postCreate-error-decode-json",
			payload:    nil,
			statusCode: http.StatusBadRequest,
		},
		{
			name: "postCreate-error-validation",
			payload: &models.PostUpdateInput{
				Title:   "",
				Content: "",
			},
			statusCode: http.StatusBadRequest,
		},
		{
			name:      "postUpdate-error-post-not-found",
			slugRoute: repository.ErrNotFoundKeyString,
			payload: &models.PostUpdateInput{
				Title:      "The updated post title",
				Content:    longText,
				CategoryId: 1,
			},
			statusCode: http.StatusNotFound,
		},
		{
			name:      "postUpdate-error-getting-post",
			slugRoute: repository.ErrUnexpectedKeyString,
			payload: &models.PostUpdateInput{
				Title:      "The updated post title",
				Content:    longText,
				CategoryId: 1,
			},
			statusCode: http.StatusInternalServerError,
		},
		{
			name:      "postUpdate-error-duplicate-title-or-post",
			slugRoute: "post-title",
			payload: &models.PostUpdateInput{
				Title:      repository.ErrDuplicateKeyString,
				Content:    longText,
				CategoryId: 1,
			},
			statusCode: http.StatusConflict,
		},
		{
			name:      "postUpdate-error-updating-post",
			slugRoute: "post-title",
			payload: &models.PostUpdateInput{
				Title:      repository.ErrUnexpectedKeyString,
				Content:    longText,
				CategoryId: 1,
			},
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		var r *http.Request
		if tt.payload == nil {
			r = httptest.NewRequest("PATCH", "/posts/{slug}", nil)
		} else {
			jsonBytes, _ := json.Marshal(tt.payload)
			r = httptest.NewRequest("PATCH", "/posts/{slug}", bytes.NewBuffer(jsonBytes))
		}

		ctx := getCtxWithParam(r, params{"slug": tt.slugRoute})
		r = r.WithContext(ctx)
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(h.post.Update)
		handler.ServeHTTP(w, r)

		if w.Code != tt.statusCode {
			t.Errorf("%s returned response code of %d, wanted %d", tt.name, w.Code, tt.statusCode)
		}
	}
}
func TestPost_Delete(t *testing.T) {
	var tests = []struct {
		name       string
		slugRoute  string
		statusCode int
	}{
		{
			name:       "postDelete-ok",
			slugRoute:  "post-title",
			statusCode: http.StatusOK,
		},
		{
			name:       "postDelete-error-post-not-found",
			slugRoute:  repository.ErrNotFoundKeyString,
			statusCode: http.StatusNotFound,
		},
		{
			name:       "postDelete-error-getting-post",
			slugRoute:  repository.ErrUnexpectedKeyString,
			statusCode: http.StatusInternalServerError,
		},
		{
			name:       "postDelete-error-deleting-post",
			slugRoute:  "get-invalid-post",
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {

		r := httptest.NewRequest("DELETE", "/posts/{slug}", nil)
		ctx := getCtxWithParam(r, params{"slug": tt.slugRoute})
		r = r.WithContext(ctx)
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(h.post.Delete)
		handler.ServeHTTP(w, r)

		if w.Code != tt.statusCode {
			t.Errorf("%s returned response code of %d, wanted %d", tt.name, w.Code, tt.statusCode)
		}
	}
}

func TestPost_GetCategories(t *testing.T) {
	var tests = []struct {
		name       string
		statusCode int
	}{
		{
			name:       "postGetCategories-ok",
			statusCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		r := httptest.NewRequest("GET", "/posts/categories", nil)
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(h.post.GetCategories)
		handler.ServeHTTP(w, r)

		if w.Code != tt.statusCode {
			t.Errorf("%s returned response code of %d, wanted %d", tt.name, w.Code, tt.statusCode)
		}
	}
}
