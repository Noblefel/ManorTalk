package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/Noblefel/ManorTalk/backend/internal/config"
	"github.com/Noblefel/ManorTalk/backend/internal/database"
	"github.com/Noblefel/ManorTalk/backend/internal/repository/postgres"
	"github.com/Noblefel/ManorTalk/backend/internal/repository/redis"
	"github.com/Noblefel/ManorTalk/backend/internal/router"
	"github.com/Noblefel/ManorTalk/backend/internal/service/auth"
	"github.com/Noblefel/ManorTalk/backend/internal/service/post"
	"github.com/Noblefel/ManorTalk/backend/internal/service/user"
	"github.com/joho/godotenv"
)

func main() {
	// If not in production, the application loads the local .env file.
	prod := flag.Bool("production", true, "Run in production mode")
	flag.Parse()

	if !*prod {
		if err := godotenv.Load(); err != nil {
			log.Fatal(err)
		}
	}

	c := config.Default().WithProductionMode(*prod)

	db, err := database.Connect(c)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Sql.Close()
	defer db.Redis.Close()

	userRepo := postgres.NewUserRepo(db)
	postRepo := postgres.NewPostRepo(db)
	cacheRepo := redis.NewRepo(db)

	authService := auth.NewAuthService(c, cacheRepo, userRepo)
	userService := user.NewUserService(c, cacheRepo, userRepo)
	postService := post.NewPostService(c, cacheRepo, postRepo)

	router := router.NewRouter(c, authService, userService, postService)

	server := &http.Server{
		Addr:    fmt.Sprint("localhost:", c.Port),
		Handler: router.Routes(),
	}

	log.Println("Starting server at port:", c.Port)

	if err = server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
