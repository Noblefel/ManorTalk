package auth

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

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

		return user, "", fmt.Errorf("%s: %w", "Error getting user by id", err)
	}

	accessToken, err := token.Generate(token.Details{
		SecretKey: s.c.AccessTokenKey,
		UserId:    user.Id,
		Duration:  s.c.AccessTokenExp,
	})

	if err != nil {
		return user, "", fmt.Errorf("%s: %w", "Error generating access token", err)
	}

	user.Password = ""

	return user, accessToken, nil
}

func (s *mockAuthService) Refresh(refreshToken string) (models.User, string, error) {
	var user models.User

	if refreshToken == ErrUnauthorized.Error() {
		return user, "", ErrUnauthorized
	}

	if refreshToken == ErrNoUser.Error() {
		return user, "", ErrNoUser
	}

	if refreshToken == http.StatusText(http.StatusInternalServerError) {
		return user, "", errors.New("unexpected error")
	}

	return user, "", nil
}
