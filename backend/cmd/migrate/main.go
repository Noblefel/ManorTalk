package main

import (
	"errors"
	"flag"
	"log"
	"os"

	"github.com/Noblefel/ManorTalk/backend/internal/config"
	"github.com/Noblefel/ManorTalk/backend/internal/database"
	"github.com/joho/godotenv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	if len(os.Args) < 2 {
		log.Println(`Missing argument. Use "up" or "down"`)
		return
	}

	// If not in production, the application loads the local .env file.
	prod := flag.Bool("production", true, "Run in production mode")
	flag.Parse()

	if !*prod {
		if err := godotenv.Load(); err != nil {
			log.Fatal(err)
		}
	}

	config := config.Default()

	db, err := database.Connect(config)
	if err != nil {
		log.Fatal("Error connecting to the database\n", err)
	}
	defer db.Sql.Close()
	defer db.Redis.Close()

	driver, err := postgres.WithInstance(db.Sql, &postgres.Config{})
	if err != nil {
		log.Fatal("Error setting up driver instance\n", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		log.Fatal("Error setting up migration\n", err)
	}

	version, dirty, err := m.Version()
	if err != nil && !errors.Is(migrate.ErrNilVersion, err) {
		log.Fatal("Error checking migration version", err)
	}

	if dirty {
		if err = m.Force(int(version)); err != nil {
			log.Fatal("Error when fixing the migration version", err)
		}
	}

	switch os.Args[1] {
	case "up":
		if err := m.Up(); err != nil {
			log.Fatal("Error when applying up migrations\n", err)
		}

		log.Println("Up Migration Success")
	case "down":
		if err := m.Down(); err != nil {
			log.Fatal("Error when applying down migrations\n", err)
		}

		log.Println("Down Migration Success")
	default:
		log.Printf(`Invalid command: %s. Use "up" or "down"`, os.Args[1])
	}
}
