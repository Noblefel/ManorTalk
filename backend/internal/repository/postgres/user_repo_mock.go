package postgres

import (
	"database/sql"
	"errors"

	"github.com/Noblefel/ManorTalk/backend/internal/models"
	"github.com/Noblefel/ManorTalk/backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

func (r *mockUserRepo) CreateUser(email, password string) (models.User, error) {
	var user models.User
	if email == repository.ErrDuplicateKeyString {
		return user, errors.New("duplicate key value")
	}

	if email == repository.ErrUnexpectedKeyString {
		return user, errors.New("unexpected error")
	}

	return user, nil
}

func (r *mockUserRepo) GetUserById(id int) (models.User, error) {
	var user models.User

	if id == repository.ErrNotFoundKeyInt {
		return user, sql.ErrNoRows
	}

	if id == repository.ErrUnexpectedKeyInt {
		return user, errors.New("unexpected error")
	}

	return user, nil
}

func (r *mockUserRepo) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	pw, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	user.Password = string(pw)

	if email == "get-invalid-user" {
		user.Id = repository.ErrUnexpectedKeyInt
		return user, nil
	}

	if email == repository.ErrNotFoundKeyString {
		return user, sql.ErrNoRows
	}

	if email == repository.ErrUnexpectedKeyString {
		return user, errors.New("unexpected error")
	}

	return user, nil
}
