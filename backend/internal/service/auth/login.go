package auth

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Noblefel/ManorTalk/backend/internal/models"
	"github.com/Noblefel/ManorTalk/backend/internal/utils/token"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (s *authService) Login(payload models.UserLoginInput) (models.User, string, string, error) {
	user, err := s.userRepo.GetUserByEmail(payload.Email)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return user, "", "", ErrNoUser
		}

		return user, "", "", fmt.Errorf("%s: %w", "Error getting user by email", err)
	}

	accessTD := token.Details{
		UserId:    user.Id,
		SecretKey: s.c.AccessTokenKey,
		Duration:  s.c.AccessTokenExp,
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return user, "", "", ErrInvalidCredentials
	}

	accessToken, err := token.Generate(accessTD)
	if err != nil {
		return user, "", "", fmt.Errorf("%s: %w", "Error generating access token", err)
	}

	refreshTD := token.Details{
		UserId:    user.Id,
		SecretKey: s.c.RefreshTokenKey,
		UniqueId:  uuid.NewString(),
		Duration:  s.c.RefreshTokenExp,
	}

	refreshToken, err := token.Generate(refreshTD)
	if err != nil {
		return user, "", "", fmt.Errorf("%s: %w", "Error generating refresh token", err)
	}

	if err = s.cacheRepo.SetRefreshToken(refreshTD); err != nil {
		return user, "", "", fmt.Errorf("%s: %w", "Error caching refresh token", err)
	}

	user.Password = ""

	return user, accessToken, refreshToken, nil
}

func (s *mockAuthService) Login(payload models.UserLoginInput) (models.User, string, string, error) {
	var user models.User

	switch payload.Password {
	case ErrNoUser.Error():
		return user, "", "", ErrNoUser
	case ErrInvalidCredentials.Error():
		return user, "", "", ErrInvalidCredentials
	case "unexpected error":
		return user, "", "", errors.New("unexpected error")
	default:
		return user, "", "", nil
	}
}
