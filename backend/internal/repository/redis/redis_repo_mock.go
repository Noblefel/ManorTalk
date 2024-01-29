package redis

import (
	"errors"

	"github.com/Noblefel/ManorTalk/backend/internal/repository"
	"github.com/Noblefel/ManorTalk/backend/internal/utils/token"
)

func (r *mockRedisRepo) SetRefreshToken(td token.Details) error {
	if td.UserId == repository.ErrUnexpectedKeyInt {
		return errors.New("Some error")
	}

	return nil
}

func (r *mockRedisRepo) GetRefreshToken(td token.Details) (string, error) {
	if td.UniqueId == repository.ErrIncorrectKey {
		return "", errors.New("Some error")
	}

	return "uuid", nil
}
