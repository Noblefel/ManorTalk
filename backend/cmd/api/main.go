package main

import (
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
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	c := config.Default()

	db, err := database.Connect(c)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Sql.Close()
	defer db.Redis.Close()

	log.Println("Starting server at port:", c.Port)

	userRepo := postgres.NewUserRepo(db)
	postRepo := postgres.NewPostRepo(db)
	cacheRepo := redis.NewRepo(db)

	authService := auth.NewAuthService(c, cacheRepo, userRepo)
	postService := post.NewPostService(c, cacheRepo, postRepo)

	router := router.NewRouter(c, db, authService, postService)

	server := &http.Server{
		Addr:    fmt.Sprint("localhost:", c.Port),
		Handler: router.Routes(),
	}

	if err = server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
