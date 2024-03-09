package auth

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Noblefel/ManorTalk/backend/internal/models"
	"github.com/Noblefel/ManorTalk/backend/internal/utils/token"
)

func (s *authService) Refresh(refreshToken string) (models.User, string, error) {
	var user models.User

	tokenDetails, err := token.Parse(s.c.RefreshTokenKey, refreshToken)
	if err != nil {
		return user, "", ErrUnauthorized
	}

	uuid, err := s.cacheRepo.GetRefreshToken(*tokenDetails)
	if err != nil || uuid != tokenDetails.UniqueId {
		return user, "", ErrUnauthorized
	}

	user, err = s.userRepo.GetUserById(tokenDetails.UserId)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return user, "", ErrNoUser
		}

		return user, "", fmt.Errorf("getting user by id: %w", err)
	}

	accessToken, err := token.Generate(token.Details{
		SecretKey: s.c.AccessTokenKey,
		UserId:    user.Id,
		Duration:  s.c.AccessTokenExp,
	})

	if err != nil {
		return user, "", fmt.Errorf("generating access token: %w", err)
	}

	user.Password = ""

	return user, accessToken, nil
}

func (s *mockAuthService) Refresh(refreshToken string) (models.User, string, error) {
	var user models.User

	switch refreshToken {
	case ErrUnauthorized.Error():
		return user, "", ErrUnauthorized
	case ErrNoUser.Error():
		return user, "", ErrNoUser
	case "unexpected error":
		return user, "", errors.New("unexpected error")
	default:
		return user, "", nil
	}
}
