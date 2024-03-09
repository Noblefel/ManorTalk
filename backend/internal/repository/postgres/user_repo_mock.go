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

func (r *mockUserRepo) GetUserById(id int) (models.User, error) {
	var user models.User

	if id == repository.NotFoundKeyInt {
		return user, sql.ErrNoRows
	}

	if id == repository.UnexpectedKeyInt {
		return user, errors.New("unexpected error")
	}

	return user, nil
}

func (r *mockUserRepo) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	pw, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	user.Password = string(pw)

	if email == "get-invalid-user" {
		user.Id = repository.UnexpectedKeyInt
		return user, nil
	}

	if email == repository.NotFoundKey {
		return user, sql.ErrNoRows
	}

	if email == repository.UnexpectedKey {
		return user, errors.New("unexpected error")
	}

	return user, nil
}

func (r *mockUserRepo) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	pw, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	user.Password = string(pw)

	if username == "get-invalid-user" {
		user.Id = repository.UnexpectedKeyInt
		return user, nil
	}

	if username == repository.NotFoundKey {
		return user, sql.ErrNoRows
	}

	if username == repository.UnexpectedKey {
		return user, errors.New("unexpected error")
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
