package redis

import (
	"errors"

	"github.com/Noblefel/ManorTalk/backend/internal/utils/token"
)

func (r *testRedisRepo) SetRefreshToken(td token.Details) error {
	if td.UserId <= -1 {
		return errors.New("Some error")
	}

	return nil
}

func (r *testRedisRepo) GetRefreshToken(td token.Details) (string, error) {
	if td.UniqueId == "incorrect" {
		return "", errors.New("Some error")
	}

	return "uuid", nil
}
