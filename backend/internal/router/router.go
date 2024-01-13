package router

import (
	"net/http"

	"github.com/Noblefel/ManorTalk/backend/internal/database"
	"github.com/Noblefel/ManorTalk/backend/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type router struct{}

func NewRouter(db *database.DB) *router {
	return &router{}
}

func (r *router) Routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.AllowContentType("application/json"))

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	mux.NotFound(handlers.NotFound)
	mux.MethodNotAllowed(handlers.MethodNotAllowed)

	return mux
}
