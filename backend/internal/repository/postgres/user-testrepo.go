package postgres

import (
	"database/sql"
	"errors"

	"github.com/Noblefel/ManorTalk/backend/internal/models"
	"github.com/Noblefel/ManorTalk/backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

func (r *testUserRepo) CreateUser(email, password string) (models.User, error) {
	var user models.User
	if email == repository.ErrDuplicateKeyString+"@example.com" {
		return user, errors.New("duplicate key value")
	}

	if email == repository.ErrUnexpectedKeyString+"@example.com" {
		return user, errors.New("unexpected error")
	}

	return user, nil
}

func (r *testUserRepo) Register(payload models.UserRegisterInput) (models.User, error) {
	var user models.User

	user, err := r.CreateUser(payload.Email, payload.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *testUserRepo) Login(payload models.UserLoginInput) (models.User, error) {
	var user models.User

	if payload.Password == repository.ErrIncorrectKey {
		return user, bcrypt.ErrMismatchedHashAndPassword
	}

	user, err := r.GetUserByEmail(payload.Email)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *testUserRepo) GetUserById(id int) (models.User, error) {
	var user models.User

	if id == repository.ErrNotFoundKeyInt {
		return user, sql.ErrNoRows
	}

	if id == repository.ErrUnexpectedKeyInt {
		return user, errors.New("unexpected error")
	}

	return user, nil
}

func (r *testUserRepo) GetUserByEmail(email string) (models.User, error) {
	var user models.User

	if email == "get-invalid-user@example.com" {
		user.Id = repository.ErrUnexpectedKeyInt
		return user, nil
	}

	if email == repository.ErrNotFoundKeyString+"@example.com" {
		return user, sql.ErrNoRows
	}

	if email == repository.ErrUnexpectedKeyString+"@example.com" {
		return user, errors.New("unexpected error")
	}

	return user, nil
}
