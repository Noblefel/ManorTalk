package auth

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Noblefel/ManorTalk/backend/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (s *authService) Register(payload models.UserRegisterInput) error {
	pw, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("%s: %w", "Error hashing password", err)
	}

	_, err = s.userRepo.CreateUser(payload.Username, payload.Email, string(pw))
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return ErrDuplicateEmail // Could also be ErrDuplicateUsername
		}

		return fmt.Errorf("%s: %w", "Error creating user", err)
	}

	return nil
}

func (s *mockAuthService) Register(payload models.UserRegisterInput) error {
	switch payload.Password {
	case ErrDuplicateEmail.Error():
		return ErrDuplicateEmail
	case "unexpected error":
		return errors.New("unexpected error")
	default:
		return nil
	}
}
