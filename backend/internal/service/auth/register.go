package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Noblefel/ManorTalk/backend/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (s *authService) Register(payload models.UserRegisterInput) (models.User, error) {
	var user models.User

	pw, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, fmt.Errorf("%s: %w", "Error hashing password", err)
	}

	user, err = s.userRepo.CreateUser(payload.Email, string(pw))
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return user, ErrDuplicateEmail
		}

		return user, fmt.Errorf("%s: %w", "Error creating user", err)
	}

	return user, nil
}

func (s *mockAuthService) Register(payload models.UserRegisterInput) (models.User, error) {
	var user models.User
	if payload.Password == ErrDuplicateEmail.Error() {
		return user, ErrDuplicateEmail
	}

	if payload.Password == http.StatusText(http.StatusInternalServerError) {
		return user, errors.New("unexpected error")
	}

	return user, nil
}
