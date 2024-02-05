package auth

import (
	"errors"
	"fmt"
	"log"

	"github.com/Noblefel/ManorTalk/backend/internal/utils/token"
)

func (s *authService) Logout(refreshToken string) error {
	tokenDetails, err := token.Parse(s.c.RefreshTokenKey, refreshToken)
	if err != nil {
		return ErrUnauthorized
	}

	uuid, err := s.cacheRepo.GetRefreshToken(*tokenDetails)
	if err != nil || uuid != tokenDetails.UniqueId {
		return ErrUnauthorized
	}

	err = s.cacheRepo.DelRefreshToken(*tokenDetails)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("%s: %w", "Error deleting refresh token", err)
	}

	return nil
}

func (s *mockAuthService) Logout(refreshToken string) error {
	switch refreshToken {
	case ErrUnauthorized.Error():
		return ErrUnauthorized
	case "unexpected error":
		return errors.New("unexpected error")
	default:
		return nil
	}
}
