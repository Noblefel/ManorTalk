package postgres

import (
	"database/sql"
	"errors"

	"github.com/Noblefel/ManorTalk/backend/internal/models"
)

func (r *testPostRepo) CreatePost(p models.Post) (models.Post, error) {
	var post models.Post

	if p.Title == "already-exists" {
		return post, errors.New("duplicate key value")
	}

	if p.Title == "unexpected-error" {
		return post, errors.New("some error")
	}

	return post, nil
}

func (r *testPostRepo) GetPostBySlug(slug string) (models.Post, error) {
	var post models.Post

	if slug == "not-found-error" {
		return post, sql.ErrNoRows
	}

	if slug == "unexpected-error" {
		return post, errors.New("some error")
	}

	if slug == "get-invalid-post" {
		post.Id = -1
		return post, nil
	}

	return post, nil
}

func (r *testPostRepo) GetPostsByCategory(category string) ([]models.Post, error) {
	posts := []models.Post{}

	if category == "unexpected-error" {
		return posts, errors.New("some error")
	}

	return posts, nil
}

func (r *testPostRepo) UpdatePost(p models.Post) error {

	if p.Title == "already-exists" {
		return errors.New("duplicate key value")
	}

	if p.Title == "unexpected-error" {
		return errors.New("some error")
	}

	if p.Title == "not-found-error" {
		return sql.ErrNoRows
	}

	return nil
}

func (r *testPostRepo) DeletePost(id int) error {
	if id < 0 {
		return errors.New("some error")
	}

	return nil
}

func (r *testPostRepo) GetCategories() ([]models.Category, error) {
	return nil, nil
}

func (r *testPostRepo) GetCategoryById(id int) (models.Category, error) {
	var category models.Category

	if id > 9999 {
		return category, sql.ErrNoRows
	}

	if id <= 0 {
		return category, errors.New("some error")
	}

	return category, nil
}
