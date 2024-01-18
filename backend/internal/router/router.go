package router

import (
	"net/http"

	"github.com/Noblefel/ManorTalk/backend/internal/config"
	"github.com/Noblefel/ManorTalk/backend/internal/database"
	"github.com/Noblefel/ManorTalk/backend/internal/handlers"
	"github.com/Noblefel/ManorTalk/backend/internal/middleware"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type router struct {
	m    *middleware.Middleware
	auth *handlers.AuthHandlers
	post *handlers.PostHandlers
}

func NewRouter(c *config.AppConfig, db *database.DB) *router {
	return &router{
		m:    middleware.New(c, db),
		auth: handlers.NewAuthHandlers(c, db),
		post: handlers.NewPostHandlers(c, db),
	}
}

func (r *router) Routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(chiMiddleware.Logger)
	mux.Use(chiMiddleware.RealIP)
	mux.Use(chiMiddleware.Recoverer)
	mux.Use(chiMiddleware.AllowContentType("application/json"))

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	mux.NotFound(handlers.NotFound)
	mux.MethodNotAllowed(handlers.MethodNotAllowed)

	api := chi.NewRouter()
	mux.Mount("/api/", api)

	api.Route("/auth", func(api chi.Router) {
		api.Post("/register", r.auth.Register)
		api.Post("/login", r.auth.Login)
		api.Post("/refresh", r.auth.Refresh)
	})

	api.Route("/posts", func(api chi.Router) {
		api.Put("/", r.post.Create)
		api.Get("/{slug}", r.post.Get)
		api.Patch("/{slug}", r.post.Update)
		api.Delete("/{slug}", r.post.Delete)
	})

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	return mux
}
