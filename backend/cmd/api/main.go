package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Noblefel/ManorTalk/backend/internal/config"
	"github.com/Noblefel/ManorTalk/backend/internal/database"
	"github.com/Noblefel/ManorTalk/backend/internal/router"
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

	router := router.NewRouter(c, db)

	server := &http.Server{
		Addr:    fmt.Sprint("localhost:", c.Port),
		Handler: router.Routes(),
	}

	if err = server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
