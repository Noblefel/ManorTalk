package redis

import (
	"context"
	"fmt"

	"github.com/Noblefel/ManorTalk/backend/internal/database"
	"github.com/Noblefel/ManorTalk/backend/internal/repository"
	"github.com/Noblefel/ManorTalk/backend/internal/utils/token"
)

type RedisRepo struct {
	db *database.DB
}

type mockRedisRepo struct{}

func NewRepo(db *database.DB) repository.CacheRepo {
	return &RedisRepo{
		db: db,
	}
}

func NewMockRepo() repository.CacheRepo {
	return &mockRedisRepo{}
}

func (r *RedisRepo) SetRefreshToken(td token.Details) error {

	_, err := r.db.Redis.Set(
		context.Background(),
		fmt.Sprint("refresh_token-", td.UserId),
		td.UniqueId,
		td.Duration,
	).Result()

	if err != nil {
		return err
	}

	return nil
}

func (r *RedisRepo) GetRefreshToken(td token.Details) (string, error) {
	uuid, err := r.db.Redis.Get(
		context.Background(),
		fmt.Sprint("refresh_token-", td.UserId),
	).Result()

	if err != nil {
		return "", err
	}

	return uuid, nil
}
