package database

import (
	"os"
	"testing"

	"github.com/Noblefel/ManorTalk/backend/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func TestConnect(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		os.Setenv("DB_HOST", "localhost")
		os.Setenv("DB_PORT", "5432")
		os.Setenv("DB_NAME", "manortalk")
		os.Setenv("DB_USER", "postgres")
		os.Setenv("DB_PASSWORD", "")
		os.Setenv("REDIS_HOST", "localhost")
		os.Setenv("REDIS_PORT", "6379")

		config := config.Default()

		db, err := Connect(config)
		if err != nil {
			t.Errorf("Connect() expected no error but got %v", err)
		}

		if db == nil {
			t.Errorf("Connect() wants *database.DB but got nil")
		}
	})

	t.Run("Fail-sql", func(t *testing.T) {
		os.Setenv("DB_HOST", "ervjojvojdofjsdvojf")
		os.Setenv("DB_PORT", "cofwoeijfowief")
		os.Setenv("DB_NAME", "jweoijewotcjo2j")
		os.Setenv("DB_USER", "34vt3jwveorijrvw")
		os.Setenv("DB_PASSWORD", "322093ruc0e0w")
		os.Setenv("REDIS_HOST", "x")
		os.Setenv("REDIS_PORT", "x")

		config := config.Default()

		_, err := Connect(config)
		if err == nil {
			t.Errorf("Connect() expected error but got none")
		}
	})

	t.Run("Fail-Redis", func(t *testing.T) {
		os.Setenv("DB_HOST", "localhost")
		os.Setenv("DB_PORT", "5432")
		os.Setenv("DB_NAME", "manortalk")
		os.Setenv("DB_USER", "postgres")
		os.Setenv("DB_PASSWORD", "")
		os.Setenv("REDIS_HOST", "x")
		os.Setenv("REDIS_PORT", "x")

		config := config.Default()

		_, err := Connect(config)
		if err == nil {
			t.Errorf("Connect() expected error but got none")
		}
	})
}

func TestConnectSQL(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		os.Setenv("DB_HOST", "localhost")
		os.Setenv("DB_PORT", "5432")
		os.Setenv("DB_NAME", "manortalk")
		os.Setenv("DB_USER", "postgres")
		os.Setenv("DB_PASSWORD", "")

		config := config.Default()

		_, err := connectSQL(config)
		if err != nil {
			t.Errorf("connectSQL() expected no error but got %v", err)
		}
	})

	t.Run("Fail-SQL", func(t *testing.T) {
		os.Setenv("DB_HOST", "ervjojvojdofjsdvojf")
		os.Setenv("DB_PORT", "cofwoeijfowief")
		os.Setenv("DB_NAME", "jweoijewotcjo2j")
		os.Setenv("DB_USER", "34vt3jwveorijrvw")
		os.Setenv("DB_PASSWORD", "322093ruc0e0w")

		config := config.Default()

		_, err := connectSQL(config)
		if err == nil {
			t.Errorf("connectSQL() expected error but got none")
		}
	})
}

func TestConnectRedis(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		os.Setenv("REDIS_HOST", "localhost")
		os.Setenv("REDIS_PORT", "6379")

		config := config.Default()

		_, err := connectRedis(config)
		if err != nil {
			t.Errorf("ConnectRedis() expected no error but got %v", err)
		}
	})

	t.Run("Fail", func(t *testing.T) {
		os.Setenv("REDIS_HOST", "x")
		os.Setenv("REDIS_PORT", "x")

		config := config.Default()

		_, err := connectRedis(config)
		if err == nil {
			t.Errorf("connectRedis() expected error but got none")
		}
	})
}
