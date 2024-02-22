package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/Noblefel/ManorTalk/backend/internal/config"
	"github.com/Noblefel/ManorTalk/backend/internal/database"
	"github.com/Noblefel/ManorTalk/backend/internal/models"
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
			name: "postCreate-error-no-category",
			payload: &models.PostCreateInput{
				Title:      service.ErrNoCategory.Error(),
				Content:    longText,
				CategoryId: 1,
			},
			statusCode: http.StatusNotFound,
		},
		{
			name: "postCreate-error-duplicate-title",
			payload: &models.PostCreateInput{
				Title:      service.ErrDuplicateTitle.Error(),
				Content:    longText,
				CategoryId: 1,
			},
			statusCode: http.StatusConflict,
		},
		{
			name: "postCreate-error-unexpected",
			payload: &models.PostCreateInput{
				Title:      "unexpected error",
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

		ctx := context.WithValue(r.Context(), "user_id", 1)
		r = r.WithContext(ctx)
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
			name:       "postGet-error-no-post",
			slugRoute:  service.ErrNoPost.Error(),
			statusCode: http.StatusNotFound,
		},
		{
			name:       "postGet-error-unexpected",
			slugRoute:  "unexpected error",
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

func TestPost_GetMany(t *testing.T) {
	var tests = []struct {
		name       string
		query      string
		statusCode int
	}{
		{
			name:       "postGetMany-ok",
			query:      "page=1&limit=10",
			statusCode: http.StatusOK,
		},
		{
			name:       "postGetMany-error-no-category",
			query:      slug.Make(service.ErrNoCategory.Error()) + "=1",
			statusCode: http.StatusNotFound,
		},
		{
			name:       "postGetMany-error-unexpected",
			query:      slug.Make("unexpected error") + "=1",
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		r := httptest.NewRequest("GET", "/posts?"+tt.query, nil)
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(h.post.GetMany)
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
			name:      "postUpdate-error-no-post",
			slugRoute: service.ErrNoPost.Error(),
			payload: &models.PostUpdateInput{
				Title:      "The updated post title",
				Content:    longText,
				CategoryId: 1,
			},
			statusCode: http.StatusNotFound,
		},
		{
			name:      "postUpdate-error-unauthorized",
			slugRoute: service.ErrUnauthorized.Error(),
			payload: &models.PostUpdateInput{
				Title:      "The updated post title",
				Content:    longText,
				CategoryId: 1,
			},
			statusCode: http.StatusUnauthorized,
		},
		{
			name:      "postUpdate-error-duplicate-title-or-post",
			slugRoute: service.ErrDuplicateTitle.Error(),
			payload: &models.PostUpdateInput{
				Title:      "The updated post title",
				Content:    longText,
				CategoryId: 1,
			},
			statusCode: http.StatusConflict,
		},
		{
			name:      "postUpdate-error-unexpected",
			slugRoute: "unexpected error",
			payload: &models.PostUpdateInput{
				Title:      "The updated post title",
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
		ctx = context.WithValue(ctx, "user_id", 1)
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
			name:       "postDelete-error-no-post",
			slugRoute:  service.ErrNoPost.Error(),
			statusCode: http.StatusNotFound,
		},
		{
			name:       "postDelete-error-unauthorized",
			slugRoute:  service.ErrUnauthorized.Error(),
			statusCode: http.StatusUnauthorized,
		},
		{
			name:       "postDelete-error-unexpected",
			slugRoute:  "unexpected error",
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {

		r := httptest.NewRequest("DELETE", "/posts/{slug}", nil)
		ctx := getCtxWithParam(r, params{"slug": tt.slugRoute})
		ctx = context.WithValue(ctx, "user_id", 1)
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
