package postgres

import (
	"database/sql"
	"errors"

	"github.com/Noblefel/ManorTalk/backend/internal/models"
	"github.com/Noblefel/ManorTalk/backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type mockUserRepo struct{}

func NewMockUserRepo() repository.UserRepo {
	return &mockUserRepo{}
}

func (r *mockUserRepo) CreateUser(username, email, password string) (int, error) {
	if email == repository.DuplicateKey {
		return 0, errors.New("duplicate key value")
	}

	if email == repository.UnexpectedKey {
		return 0, errors.New("unexpected error")
	}

	return 1, nil
}

func (r *mockUserRepo) GetUser(filters models.UserFilters) (models.User, error) {
	var user models.User
	pw, _ := bcrypt.GenerateFromPassword([]byte("password"), 7)
	user.Password = string(pw)

	if filters.Id == repository.NotFoundKeyInt ||
		filters.Email == repository.NotFoundKey ||
		filters.Username == repository.NotFoundKey {
		return user, sql.ErrNoRows
	}

	if filters.Id == repository.UnexpectedKeyInt ||
		filters.Email == repository.UnexpectedKey ||
		filters.Username == repository.UnexpectedKey {
		return user, errors.New("unexpected error")
	}

	if filters.Email == "get-invalid-user" || filters.Username == "get-invalid-user" {
		user.Id = repository.UnexpectedKeyInt
		return user, nil
	}

	return user, nil
}

func (r *mockUserRepo) UpdateUser(u models.User) error {

	if u.Name == repository.DuplicateKey {
		return errors.New("duplicate key value")
	}

	if u.Name == repository.UnexpectedKey {
		return errors.New("some error")
	}

	if u.Name == repository.NotFoundKey {
		return sql.ErrNoRows
	}

	return nil
}
