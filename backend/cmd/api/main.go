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

	config := config.Default()

	db, err := database.Connect(config)
	if err != nil {
		log.Fatalln("Error connecting to the database", err)
	}
	defer db.Sql.Close()

	log.Println("Starting server at port:", config.Port)

	router := router.NewRouter(db)

	server := &http.Server{
		Addr:    fmt.Sprint("localhost:", config.Port),
		Handler: router.Routes(),
	}

	if err = server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
