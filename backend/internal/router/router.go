package router

import (
	"net/http"

	"github.com/Noblefel/ManorTalk/backend/internal/config"
	"github.com/Noblefel/ManorTalk/backend/internal/handlers"
	"github.com/Noblefel/ManorTalk/backend/internal/middleware"
	"github.com/Noblefel/ManorTalk/backend/internal/service/auth"
	"github.com/Noblefel/ManorTalk/backend/internal/service/post"
	"github.com/Noblefel/ManorTalk/backend/internal/service/user"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type router struct {
	m    *middleware.Middleware
	auth *handlers.AuthHandlers
	user *handlers.UserHandlers
	post *handlers.PostHandlers
}

func NewRouter(
	c *config.AppConfig,
	as auth.AuthService,
	us user.UserService,
	ps post.PostService,
) *router {
	return &router{
		m:    middleware.New(c),
		auth: handlers.NewAuthHandlers(as),
		user: handlers.NewUserHandlers(us),
		post: handlers.NewPostHandlers(ps),
	}
}

func (r *router) Routes() http.Handler {
	mux := chi.NewRouter()

	// mux.Use(chiMiddleware.Logger)
	mux.Use(chiMiddleware.RealIP)
	mux.Use(chiMiddleware.Recoverer)
	mux.Use(chiMiddleware.AllowContentType(
		"application/json", "multipart/form-data",
	))

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.NotFound(handlers.NotFound)
	mux.MethodNotAllowed(handlers.MethodNotAllowed)

	api := chi.NewRouter()
	mux.Mount("/api/", api)

	r.authRouter(api)
	r.postRouter(api)
	r.userRouter(api)

	fileServer := http.FileServer(http.Dir("./images/"))
	mux.Handle("/images/*", http.StripPrefix("/images", fileServer))

	return mux
}

func (r *router) authRouter(api *chi.Mux) {
	api.Route("/auth", func(api chi.Router) {
		api.Post("/register", r.auth.Register)
		api.Post("/login", r.auth.Login)
		api.Post("/refresh", r.auth.Refresh)
		api.Post("/logout", r.auth.Logout)
	})
}

func (r *router) postRouter(api *chi.Mux) {
	api.Route("/posts", func(api chi.Router) {
		api.Get("/", r.post.GetMany)
		api.Get("/{slug}", r.post.Get)
		api.Get("/categories", r.post.GetCategories)

		api.Group(func(api chi.Router) {
			api.Use(r.m.Auth)
			api.Post("/", r.post.Create)
			api.Patch("/{slug}", r.post.Update)
			api.Delete("/{slug}", r.post.Delete)
		})
	})
}

func (r *router) userRouter(api *chi.Mux) {
	api.Route("/users", func(api chi.Router) {
		api.Post("/check-username", r.user.CheckUsername)
		api.Get("/{username}", r.user.Get)

		api.Group(func(api chi.Router) {
			api.Use(r.m.Auth)
			api.Patch("/{username}", r.user.UpdateProfile)
		})
	})
}
