package postgres

import (
	"database/sql"
	"errors"

	"github.com/Noblefel/ManorTalk/backend/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (r *testUserRepo) CreateUser(email, password string) (models.User, error) {
	var user models.User
	if email == "alreadyexists@error.com" {
		return user, errors.New("duplicate key value")
	}

	if email == "unexpected@error.com" {
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

	if payload.Password == "incorrectpassword" {
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

	if id > 9999 {
		return user, sql.ErrNoRows
	}

	if id <= 0 {
		return user, errors.New("unexpected error")
	}

	return user, nil
}

func (r *testUserRepo) GetUserByEmail(email string) (models.User, error) {
	var user models.User

	if email == "invaliduser@example.com" {
		user.Id = -1
		return user, nil
	}

	if email == "notfound@error.com" {
		return user, sql.ErrNoRows
	}

	if email == "unexpected@error.com" {
		return user, errors.New("unexpected error")
	}

	return user, nil
}
