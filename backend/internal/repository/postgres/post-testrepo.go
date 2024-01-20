package postgres

import (
	"database/sql"
	"errors"

	"github.com/Noblefel/ManorTalk/backend/internal/models"
	"github.com/Noblefel/ManorTalk/backend/internal/repository"
)

func (r *testPostRepo) CreatePost(p models.Post) (models.Post, error) {
	var post models.Post

	if p.Title == repository.ErrDuplicateKeyString {
		return post, errors.New("duplicate key value")
	}

	if p.Title == repository.ErrUnexpectedKeyString {
		return post, errors.New("some error")
	}

	return post, nil
}

func (r *testPostRepo) GetPostBySlug(slug string) (models.Post, error) {
	var post models.Post

	if slug == repository.ErrNotFoundKeyString {
		return post, sql.ErrNoRows
	}

	if slug == repository.ErrUnexpectedKeyString {
		return post, errors.New("some error")
	}

	if slug == "get-invalid-post" {
		post.Id = repository.ErrUnexpectedKeyInt
		return post, nil
	}

	return post, nil
}

func (r *testPostRepo) GetPostsByCategory(category string) ([]models.Post, error) {
	posts := []models.Post{}

	if category == repository.ErrUnexpectedKeyString {
		return posts, errors.New("some error")
	}

	return posts, nil
}

func (r *testPostRepo) UpdatePost(p models.Post) error {

	if p.Title == repository.ErrDuplicateKeyString {
		return errors.New("duplicate key value")
	}

	if p.Title == repository.ErrUnexpectedKeyString {
		return errors.New("some error")
	}

	if p.Title == repository.ErrNotFoundKeyString {
		return sql.ErrNoRows
	}

	return nil
}

func (r *testPostRepo) DeletePost(id int) error {
	if id == repository.ErrUnexpectedKeyInt {
		return errors.New("some error")
	}

	return nil
}

func (r *testPostRepo) GetCategories() ([]models.Category, error) {
	return nil, nil
}

func (r *testPostRepo) GetCategoryById(id int) (models.Category, error) {
	var category models.Category

	if id == repository.ErrNotFoundKeyInt {
		return category, sql.ErrNoRows
	}

	if id == repository.ErrUnexpectedKeyInt {
		return category, errors.New("some error")
	}

	return category, nil
}
