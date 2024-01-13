package config

import (
	"os"
	"time"
)

type AppConfig struct {
	InDevelopment bool
	Port          int
	DB            dbConfig
}

type dbConfig struct {
	Host, Port, Name, User, Password string
	MaxOpenConns, MaxIdleConns       int
	MaxLifetime                      time.Duration
}

func Default() *AppConfig {
	return &AppConfig{
		InDevelopment: false,
		Port:          8080,
		DB: dbConfig{
			Host:         os.Getenv("DB_HOST"),
			Port:         os.Getenv("DB_PORT"),
			Name:         os.Getenv("DB_NAME"),
			User:         os.Getenv("DB_USER"),
			Password:     os.Getenv("DB_PASSWORD"),
			MaxOpenConns: 10,
			MaxIdleConns: 5,
			MaxLifetime:  5 * time.Minute,
		},
	}
}
