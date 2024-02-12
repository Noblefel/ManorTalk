package user

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Noblefel/ManorTalk/backend/internal/models"
)

func (s *userService) Get(username string) (models.User, error) {
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return user, ErrNoUser
		}

		return user, fmt.Errorf("%s: %w", "Error getting user by username", err)
	}

	user.Email = ""
	user.Password = ""

	return user, nil
}

func (s *mockUserService) Get(username string) (models.User, error) {
	var user models.User
	switch username {
	case ErrNoUser.Error():
		return user, ErrNoUser
	case "unexpected error":
		return user, errors.New("unexpected error")
	default:
		return user, nil
	}
}
