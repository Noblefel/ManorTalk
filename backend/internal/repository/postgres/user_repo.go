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

func NewUserRepo(db *database.DB) repository.UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) CreateUser(username, email, password string) (int, error) {
	query := `
		INSERT INTO users (username, email, password, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
		`

	var id int

	err := r.db.Sql.QueryRow(query,
		username,
		email,
		password,
		time.Now(),
		time.Now(),
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserRepo) GetUserById(id int) (models.User, error) {
	var user models.User

	query := `
	SELECT 
		u.id, 
		COALESCE(u.name,''), 
		u.username, 
		COALESCE(u.avatar,''), 
		COALESCE(u.bio, ''),
		u.email, 
		u.password, 
		u.created_at, 
		u.updated_at, 
		COUNT(p.id) AS posts_count
	FROM users u 
	LEFT JOIN posts p ON (p.user_id = u.id) 
	WHERE u.id = $1
	GROUP BY u.id`

	err := r.db.Sql.QueryRow(query, id).Scan(
		&user.Id,
		&user.Name,
		&user.Username,
		&user.Avatar,
		&user.Bio,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.PostsCount,
	)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepo) GetUserByEmail(email string) (models.User, error) {
	var user models.User

	query := `
	SELECT 
		u.id, 
		COALESCE(u.name,''), 
		u.username, 
		COALESCE(u.avatar,''), 
		COALESCE(u.bio, ''),
		u.email, 
		u.password, 
		u.created_at, 
		u.updated_at, 
		COUNT(p.id) AS posts_count
	FROM users u 
	LEFT JOIN posts p ON (p.user_id = u.id) 
	WHERE u.email = $1
	GROUP BY u.id`

	err := r.db.Sql.QueryRow(query, email).Scan(
		&user.Id,
		&user.Name,
		&user.Username,
		&user.Avatar,
		&user.Bio,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.PostsCount,
	)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepo) GetUserByUsername(username string) (models.User, error) {
	var user models.User

	query := `
	SELECT 
		u.id, 
		COALESCE(u.name,''), 
		u.username, 
		COALESCE(u.avatar,''), 
		COALESCE(u.bio, ''),
		u.email, 
		u.password, 
		u.created_at, 
		u.updated_at, 
		COUNT(p.id) AS posts_count
	FROM users u 
	LEFT JOIN posts p ON (p.user_id = u.id) 
	WHERE u.username = $1
	GROUP BY u.id`

	err := r.db.Sql.QueryRow(query, username).Scan(
		&user.Id,
		&user.Name,
		&user.Username,
		&user.Avatar,
		&user.Bio,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.PostsCount,
	)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepo) UpdateUser(u models.User) error {
	query := `
	UPDATE users 
		SET 
			name = NULLIF($1, ''), 
			username = $2, 
			avatar = COALESCE(NULLIF($3, ''), avatar),
			bio = NULLIF($4, ''), 
			email = COALESCE(NULLIF($5, ''), email), 
			password = COALESCE(NULLIF($6, ''), password), 
			updated_at = $7 
	WHERE id = $8
`

	_, err := r.db.Sql.Exec(query,
		u.Name,
		u.Username,
		u.Avatar,
		u.Bio,
		u.Email,
		u.Password,
		time.Now(),
		u.Id,
	)

	if err != nil {
		return err
	}

	return nil
}
