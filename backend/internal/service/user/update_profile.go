package user

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/Noblefel/ManorTalk/backend/internal/models"
)

func (s *userService) UpdateProfile(payload models.UpdateProfileInput, username string, authId int) error {
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return ErrNoUser
		}

		return fmt.Errorf("%s: %w", "Error getting user by slug", err)
	}

	if authId != user.Id {
		return ErrUnauthorized
	}

	user.Name = payload.Name
	user.Username = payload.Username

	err = s.userRepo.UpdateUser(user)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return ErrDuplicateUsername
		}

		return fmt.Errorf("%s: %w", "Error updating user", err)
	}

	return nil
}

func (s *mockUserService) UpdateProfile(payload models.UpdateProfileInput, username string, authId int) error {
	switch username {
	case ErrNoUser.Error():
		return ErrNoUser
	case ErrUnauthorized.Error():
		return ErrUnauthorized
	case ErrDuplicateUsername.Error():
		return ErrDuplicateUsername
	case "unexpected error":
		return errors.New("unexpected error")
	default:
		return nil
	}
}
