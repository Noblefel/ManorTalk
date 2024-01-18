package postgres

import (
	"time"

	"github.com/Noblefel/ManorTalk/backend/internal/database"
	"github.com/Noblefel/ManorTalk/backend/internal/models"
	"github.com/Noblefel/ManorTalk/backend/internal/repository"
)

type PostRepo struct {
	db *database.DB
}

type testPostRepo struct{}

func NewPostRepo(db *database.DB) repository.PostRepo {
	return &PostRepo{
		db: db,
	}
}

func NewTestPostRepo() repository.PostRepo {
	return &testPostRepo{}
}

func (r *PostRepo) GetPostBySlug(slug string) (models.Post, error) {
	var post models.Post

	query := `
		SELECT id, user_id, title, slug, excerpt, content, category_id, created_at, updated_at 
		FROM posts
		WHERE slug = $1 
	`

	err := r.db.Sql.QueryRow(query, slug).Scan(
		&post.Id,
		&post.UserId,
		&post.Title,
		&post.Slug,
		&post.Excerpt,
		&post.Content,
		&post.CategoryId,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		return post, err
	}

	return post, nil
}

func (r *PostRepo) CreatePost(p models.Post) (models.Post, error) {
	var post models.Post

	query := `
		INSERT INTO posts (user_id, title, slug, excerpt, content, category_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, user_id, title, slug, excerpt, content, category_id, created_at, updated_at
	`

	err := r.db.Sql.QueryRow(query,
		p.UserId,
		p.Title,
		p.Slug,
		p.Excerpt,
		p.Content,
		p.CategoryId,
		time.Now(),
		time.Now(),
	).Scan(
		&post.Id,
		&post.UserId,
		&post.Title,
		&post.Slug,
		&post.Excerpt,
		&post.Content,
		&post.CategoryId,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		return post, err
	}

	return post, nil
}

func (r *PostRepo) UpdatePost(p models.Post) error {
	query := `
		UPDATE posts 
			SET 
				title = $1, 
				slug = $2, 
				excerpt = $3, 
				content = $4, 
				category_id = $5, 
				updated_at = $6 
		WHERE id = $7
	`

	_, err := r.db.Sql.Exec(query,
		p.Title,
		p.Slug,
		p.Excerpt,
		p.Content,
		p.CategoryId,
		time.Now(),
		p.Id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *PostRepo) DeletePost(id int) error {
	query := `DELETE FROM posts WHERE id = $1`

	_, err := r.db.Sql.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
