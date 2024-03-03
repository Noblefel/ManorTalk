package postgres

import (
	"database/sql"
	"time"

	"github.com/Noblefel/ManorTalk/backend/internal/database"
	"github.com/Noblefel/ManorTalk/backend/internal/models"
	"github.com/Noblefel/ManorTalk/backend/internal/repository"
	"github.com/Noblefel/ManorTalk/backend/internal/utils/pagination"
)

type PostRepo struct {
	db *database.DB
}

func NewPostRepo(db *database.DB) repository.PostRepo {
	return &PostRepo{
		db: db,
	}
}

func (r *PostRepo) GetPosts(pgMeta *pagination.Meta, filters models.PostsFilters) ([]models.Post, error) {
	posts := []models.Post{}

	query := `
		SELECT 
			p.id, 
			p.user_id, 
			p.title, 
			p.slug, 
			COALESCE(p.excerpt, ''),
			COALESCE(p.image, ''),  
			p.content, 
			p.category_id, 
			p.created_at, 
			p.updated_at,
			c.name, 
			c.slug,
			COALESCE(u.name, ''), 
			u.username, 
			COALESCE(u.avatar, '')
		FROM posts p
		LEFT JOIN categories c ON (p.category_id = c.id)
		LEFT JOIN users u ON (p.user_id = u.id)	
		`

	if filters.Category != "" {
		query += "WHERE c.slug = $3\n"
	} else {
		query += "WHERE c.slug != $3\n"
	}

	if filters.Search != "" {
		query += "AND p.title LIKE CONCAT('%', $4::text, '%')\n"
	} else {
		query += "AND p.title != $4\n"
	}

	if filters.UserId != 0 {
		query += "AND p.user_id = $5\n"
	} else {
		query += "AND p.user_id != $5\n"
	}

	var rows *sql.Rows
	var err error

	if filters.Cursor != 0 {
		query += "AND p.id >= $1 LIMIT $2"
		rows, err = r.db.Sql.Query(query,
			filters.Cursor,
			filters.Limit,
			filters.Category,
			filters.Search,
			filters.UserId,
		)
	} else {
		query += "OFFSET $1 LIMIT $2"
		rows, err = r.db.Sql.Query(query,
			pgMeta.Offset,
			filters.Limit,
			filters.Category,
			filters.Search,
			filters.UserId,
		)
	}

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
			&post.Image,
			&post.Content,
			&post.CategoryId,
			&post.CreatedAt,
			&post.UpdatedAt,
			&post.Category.Name,
			&post.Category.Slug,
			&post.User.Name,
			&post.User.Username,
			&post.User.Avatar,
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
			p.id, 
			p.user_id, 
			p.title, 
			p.slug, 
			COALESCE(p.excerpt, ''), 
			COALESCE(p.image, ''), 
			p.content, 
			p.category_id, 
			p.created_at, 
			p.updated_at,
			COALESCE(u.name, ''), 
			u.username, 
			COALESCE(u.avatar, ''),
			c.name, 
			c.slug 
		FROM posts p
		LEFT JOIN users u ON (p.user_id = u.id)
		LEFT JOIN categories c ON (p.category_id = c.id)
		WHERE p.slug = $1 
	`

	err := r.db.Sql.QueryRow(query, slug).Scan(
		&post.Id,
		&post.UserId,
		&post.Title,
		&post.Slug,
		&post.Excerpt,
		&post.Image,
		&post.Content,
		&post.CategoryId,
		&post.CreatedAt,
		&post.UpdatedAt,
		&post.User.Name,
		&post.User.Username,
		&post.User.Avatar,
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
		INSERT INTO posts (
			user_id, 
			title, 
			slug, 
			excerpt,
			image, 
			content, 
			category_id, 
			created_at, 
			updated_at
		)
		VALUES ($1, $2, $3, $4, NULLIF($5, ''), $6, $7, $8, $9)
		RETURNING 
			id, 
			user_id, 
			title, 
			slug, 
			excerpt,
			COALESCE(image, ''), 
			content, 
			category_id, 
			created_at, 
			updated_at
	`

	err := r.db.Sql.QueryRow(query,
		p.UserId,
		p.Title,
		p.Slug,
		p.Excerpt,
		p.Image,
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
		&post.Image,
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
				image = COALESCE(NULLIF($4, ''), image),
				content = $5, 
				category_id = $6, 
				updated_at = $7 
		WHERE id = $8
	`

	_, err := r.db.Sql.Exec(query,
		p.Title,
		p.Slug,
		p.Excerpt,
		p.Image,
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

	if filters.Search != "" {
		query += "AND p.title LIKE CONCAT('%', $2::text, '%')\n"
	} else {
		query += "AND p.title != $2\n"
	}

	err := r.db.Sql.QueryRow(query, filters.Category, filters.Search).Scan(&count)
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
