package user

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

func (s *userService) CheckUsername(username string) error {
	_, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil
		}

		return fmt.Errorf("%s: %w", "Error getting user by username", err)
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