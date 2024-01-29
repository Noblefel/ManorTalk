package postgres

import (
	"time"

	"github.com/Noblefel/ManorTalk/backend/internal/database"
	"github.com/Noblefel/ManorTalk/backend/internal/models"
	"github.com/Noblefel/ManorTalk/backend/internal/repository"
)

type UserRepo struct {
	db *database.DB
}

type mockUserRepo struct{}

func NewUserRepo(db *database.DB) repository.UserRepo {
	return &UserRepo{
		db: db,
	}
}

func NewMockUserRepo() repository.UserRepo {
	return &mockUserRepo{}
}

func (r *UserRepo) CreateUser(email, password string) (models.User, error) {
	query := `
		INSERT INTO users (email, password, created_at, updated_at) 
		VALUES ($1, $2, $3, $4)
		RETURNING id, email, created_at
		`

	var user models.User

	err := r.db.Sql.QueryRow(query,
		email,
		password,
		time.Now(),
		time.Now(),
	).Scan(&user.Id, &user.Email, &user.CreatedAt)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepo) GetUserById(id int) (models.User, error) {
	var user models.User

	query := `
	SELECT id, email, password, created_at, updated_at 
	FROM users 
	WHERE id = $1`

	err := r.db.Sql.QueryRow(query, id).Scan(
		&user.Id,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepo) GetUserByEmail(email string) (models.User, error) {
	var user models.User

	query := `
	SELECT id, email, password, created_at, updated_at 
	FROM users 
	WHERE email = $1`

	err := r.db.Sql.QueryRow(query, email).Scan(
		&user.Id,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return user, err
	}

	return user, nil
}