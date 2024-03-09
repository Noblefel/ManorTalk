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
		return fmt.Errorf("hashing password: %w", err)
	}

	_, err = s.userRepo.CreateUser(payload.Username, payload.Email, string(pw))
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return ErrDuplicateEmail // or ErrDuplicateUsername
		}

		return fmt.Errorf("creating user: %w", err)
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
