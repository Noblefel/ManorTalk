package postgres

import (
	"time"

	"github.com/Noblefel/ManorTalk/backend/internal/database"
	"github.com/Noblefel/ManorTalk/backend/internal/models"
	"github.com/Noblefel/ManorTalk/backend/internal/repository"
	"github.com/Noblefel/ManorTalk/backend/internal/utils/pagination"
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

func (r *PostRepo) GetPosts(pgMeta *pagination.Meta, filters models.PostsFilters) ([]models.Post, error) {
	posts := []models.Post{}

	query := `
		SELECT 
			p.id, p.user_id, p.title, p.slug, COALESCE(p.excerpt, ''), p.content, p.category_id, p.created_at, p.updated_at,
			c.name, c.slug 
		FROM posts p
		LEFT JOIN categories c ON (p.category_id = c.id)`

	if filters.Category != "" {
		query += "WHERE c.slug = $3\n"
	} else {
		query += "WHERE c.slug != $3\n"
	}

	if filters.Order == "desc" {
		query += "ORDER BY p.id DESC\n"
	} else {
		query += "ORDER BY p.id ASC\n"
	}

	query += "OFFSET $1 LIMIT $2"

	rows, err := r.db.Sql.Query(query,
		pgMeta.Offset,
		pgMeta.Limit,
		filters.Category,
	)
	defer rows.Close()

	if err != nil {
		return posts, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post

		err = rows.Scan(
			&post.Id,
			&post.UserId,
			&post.Title,
			&post.Slug,
			&post.Excerpt,
			&post.Content,
			&post.CategoryId,
			&post.CreatedAt,
			&post.UpdatedAt,
			&post.Category.Name,
			&post.Category.Slug,
		)

		if err != nil {
			return posts, err
		}

		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return posts, err
	}

	return posts, nil
}

func (r *PostRepo) GetPostBySlug(slug string) (models.Post, error) {
	var post models.Post

	query := `
		SELECT 
			p.id, p.user_id, p.title, p.slug, COALESCE(p.excerpt, ''), p.content, p.category_id, p.created_at, p.updated_at,
			c.name, c.slug 
		FROM posts p
		LEFT JOIN categories c ON (p.category_id = c.id)
		WHERE p.slug = $1 
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
		&post.Category.Name,
		&post.Category.Slug,
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

func (r *PostRepo) CountPosts(filters models.PostsFilters) (int, error) {
	var count int
	query := `		
	SELECT 
		COUNT(*)
	FROM posts p
	LEFT JOIN categories c ON (p.category_id = c.id)
	`

	if filters.Category != "" {
		query += "WHERE c.slug = $1\n"
	} else {
		query += "WHERE c.slug != $1\n"
	}

	err := r.db.Sql.QueryRow(query, filters.Category).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *PostRepo) GetCategories() ([]models.Category, error) {
	categories := []models.Category{}

	query := `SELECT id, name, slug FROM categories`

	rows, err := r.db.Sql.Query(query)
	if err != nil {
		return categories, err
	}
	defer rows.Close()

	for rows.Next() {
		var c models.Category

		err = rows.Scan(
			&c.Id,
			&c.Name,
			&c.Slug,
		)

		if err != nil {
			return categories, err
		}

		categories = append(categories, c)
	}

	if err = rows.Err(); err != nil {
		return categories, err
	}

	return categories, nil
}

func (r *PostRepo) GetCategoryById(id int) (models.Category, error) {
	var category models.Category

	query := `SELECT id, name, slug FROM categories WHERE id = $1`

	err := r.db.Sql.QueryRow(query, id).Scan(
		&category.Id,
		&category.Name,
		&category.Slug,
	)

	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *PostRepo) GetCategoryBySlug(slug string) (models.Category, error) {
	var category models.Category

	query := `SELECT id, name, slug FROM categories WHERE slug = $1`

	err := r.db.Sql.QueryRow(query, slug).Scan(
		&category.Id,
		&category.Name,
		&category.Slug,
	)

	if err != nil {
		return category, err
	}

	return category, nil
}
