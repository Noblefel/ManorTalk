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

	// mux.Use(chiMiddleware.Logger)
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

	r.authRouter(api)
	r.postRouter(api)

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	return mux
}

func (r *router) authRouter(api *chi.Mux) {
	api.Route("/auth", func(api chi.Router) {
		api.Post("/register", r.auth.Register)
		api.Post("/login", r.auth.Login)
		api.Post("/refresh", r.auth.Refresh)
	})
}

func (r *router) postRouter(api *chi.Mux) {
	api.Route("/posts", func(api chi.Router) {
		api.Get("/", r.post.GetMany)
		api.Get("/{slug}", r.post.Get)
		api.Get("/categories", r.post.GetCategories)
		// api.Get("/c/{category}", r.post.GetByCategory)

		api.Group(func(api chi.Router) {
			api.Use(r.m.Auth)
			api.Post("/", r.post.Create)
			api.Patch("/{slug}", r.post.Update)
			api.Delete("/{slug}", r.post.Delete)
		})
	})
}
