package config

import (
	"os"
	"strconv"
	"time"
)

type AppConfig struct {
	InProduction    bool
	Port            int
	AccessTokenKey  string
	AccessTokenExp  time.Duration
	RefreshTokenKey string
	RefreshTokenExp time.Duration
	DB              dbConfig
}

type dbConfig struct {
	Host, Name, User, Password, RedisHost       string
	Port, RedisPort, MaxOpenConns, MaxIdleConns int
	MaxLifetime                                 time.Duration
}

func Default() *AppConfig {
	port, _ := strconv.Atoi(os.Getenv("APPLICATION_PORT"))
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	redisPort, _ := strconv.Atoi(os.Getenv("REDIS_PORT"))

	return &AppConfig{
		InProduction:    false,
		Port:            port,
		AccessTokenKey:  os.Getenv("ACCESS_TOKEN_KEY"),
		AccessTokenExp:  time.Duration(15 * time.Minute),
		RefreshTokenKey: os.Getenv("REFRESH_TOKEN_KEY"),
		RefreshTokenExp: time.Duration(240 * time.Hour),
		DB: dbConfig{
			Host:         os.Getenv("DB_HOST"),
			Port:         dbPort,
			Name:         os.Getenv("DB_NAME"),
			User:         os.Getenv("DB_USER"),
			Password:     os.Getenv("DB_PASSWORD"),
			RedisHost:    os.Getenv("REDIS_HOST"),
			RedisPort:    redisPort,
			MaxOpenConns: 10,
			MaxIdleConns: 5,
			MaxLifetime:  5 * time.Minute,
		},
	}
}

func (c *AppConfig) WithProductionMode(b bool) *AppConfig {
	c.InProduction = b
	return c
}
