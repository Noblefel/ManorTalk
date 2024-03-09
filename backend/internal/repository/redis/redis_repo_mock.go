package redis

import (
	"errors"

	"github.com/Noblefel/ManorTalk/backend/internal/repository"
	"github.com/Noblefel/ManorTalk/backend/internal/utils/token"
)

type mockRedisRepo struct{}

func NewMockRepo() repository.CacheRepo {
	return &mockRedisRepo{}
}

func (r *mockRedisRepo) SetRefreshToken(td token.Details) error {
	if td.UserId == repository.UnexpectedKeyInt {
		return errors.New("Some error")
	}

	return nil
}

func (r *mockRedisRepo) GetRefreshToken(td token.Details) (string, error) {
	if td.UniqueId == repository.IncorrectKey {
		return "", errors.New("Some error")
	}

	return "uuid", nil
}

func (r *mockRedisRepo) DelRefreshToken(td token.Details) error {
	if td.UserId == repository.UnexpectedKeyInt {
		return errors.New("Some error")
	}

	return nil
}
