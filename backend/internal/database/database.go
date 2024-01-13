package database

import (
	"database/sql"
	"fmt"

	"github.com/Noblefel/ManorTalk/backend/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type DB struct {
	Sql *sql.DB
}

func Connect(c *config.AppConfig) (*DB, error) {
	var db DB

	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s",
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

	db.Sql = conn

	return &db, nil
}
