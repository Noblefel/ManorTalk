package user

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Noblefel/ManorTalk/backend/internal/models"
	"github.com/gosimple/slug"
)

func (s *userService) CheckUsername(username string) error {
	_, err := s.userRepo.GetUser(models.UserFilters{Username: username})
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil
		}

		return fmt.Errorf("getting user by username: %w", err)
	}

	return ErrDuplicateUsername
}

func (s *mockUserService) CheckUsername(username string) error {
	switch username {
	case slug.Make(ErrDuplicateUsername.Error()):
		return ErrDuplicateUsername
	case slug.Make("unexpected error"):
		return errors.New("unexpected error")
	default:
		return nil
	}
}
