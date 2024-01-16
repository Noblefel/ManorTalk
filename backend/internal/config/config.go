package config

import (
	"os"
	"time"
)

type AppConfig struct {
	InDevelopment   bool
	Port            int
	AccessTokenKey  string
	AccessTokenExp  time.Duration
	RefreshTokenKey string
	RefreshTokenExp time.Duration
	DB              dbConfig
}

type dbConfig struct {
	Host, Port, Name, User, Password, RedisHost, RedisPort string
	MaxOpenConns, MaxIdleConns                             int
	MaxLifetime                                            time.Duration
}

func Default() *AppConfig {
	accessTokenExp, _ := time.ParseDuration(os.Getenv("ACCESS_TOKEN_EXP"))
	refreshTokenExp, _ := time.ParseDuration(os.Getenv("REFRESH_TOKEN_EXP"))

	return &AppConfig{
		InDevelopment:   false,
		Port:            8080,
		AccessTokenKey:  os.Getenv("ACCESS_TOKEN_KEY"),
		AccessTokenExp:  accessTokenExp,
		RefreshTokenKey: os.Getenv("JWT_REFRESH_KEY"),
		RefreshTokenExp: refreshTokenExp,
		DB: dbConfig{
			Host:         os.Getenv("DB_HOST"),
			Port:         os.Getenv("DB_PORT"),
			Name:         os.Getenv("DB_NAME"),
			User:         os.Getenv("DB_USER"),
			Password:     os.Getenv("DB_PASSWORD"),
			RedisHost:    os.Getenv("REDIS_HOST"),
			RedisPort:    os.Getenv("REDIS_PORT"),
			MaxOpenConns: 10,
			MaxIdleConns: 5,
			MaxLifetime:  5 * time.Minute,
		},
	}
}
