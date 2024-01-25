package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/Noblefel/ManorTalk/backend/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/redis/go-redis/v9"
)

type DB struct {
	Sql   *sql.DB
	Redis *redis.Client
}

func Connect(c *config.AppConfig) (*DB, error) {
	sql, err := connectSQL(c)
	if err != nil {
		log.Println("ERROR connecting to SQL")
		return nil, err
	}

	redis, err := connectRedis(c)
	if err != nil {
		log.Println("ERROR connecting to Redis")
		return nil, err
	}

	return &DB{
		Sql:   sql,
		Redis: redis,
	}, nil
}

func connectSQL(c *config.AppConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s",
		c.DB.Host,
		c.DB.Port,
		c.DB.Name,
		c.DB.User,
		c.DB.Password,
	)

	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = conn.Ping(); err != nil {
		return nil, err
	}

	conn.SetMaxOpenConns(c.DB.MaxOpenConns)
	conn.SetMaxIdleConns(c.DB.MaxIdleConns)
	conn.SetConnMaxLifetime(c.DB.MaxLifetime)

	return conn, nil
}

func connectRedis(c *config.AppConfig) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprint(c.DB.RedisHost, ":", c.DB.RedisPort),
	})

	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}

	return rdb, nil
}
