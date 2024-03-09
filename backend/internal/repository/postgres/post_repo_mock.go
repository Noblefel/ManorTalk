package postgres

import (
	"database/sql"
	"errors"

	"github.com/Noblefel/ManorTalk/backend/internal/models"
	"github.com/Noblefel/ManorTalk/backend/internal/repository"
	"github.com/Noblefel/ManorTalk/backend/internal/utils/pagination"
)

type mockPostRepo struct{}

func NewMockPostRepo() repository.PostRepo {
	return &mockPostRepo{}
}

func (r *mockPostRepo) CreatePost(p models.Post) (models.Post, error) {
	var post models.Post

	if p.Title == repository.DuplicateKey {
		return post, errors.New("duplicate key value")
	}

	if p.Title == repository.UnexpectedKey {
		return post, errors.New("some error")
	}

	return post, nil
}

func (r *mockPostRepo) GetPosts(pgMeta *pagination.Meta, filters models.PostsFilters) ([]models.Post, error) {
	posts := []models.Post{}

	if filters.Order == repository.UnexpectedKey {
		return posts, errors.New("some error")
	}

	return posts, nil
}

func (r *mockPostRepo) GetPostBySlug(slug string) (models.Post, error) {
	var post models.Post

	if slug == repository.NotFoundKey {
		return post, sql.ErrNoRows
	}

	if slug == repository.UnexpectedKey {
		return post, errors.New("some error")
	}

	if slug == "get-invalid-post" {
		post.Id = repository.UnexpectedKeyInt
		return post, nil
	}

	return post, nil
}

func (r *mockPostRepo) UpdatePost(p models.Post) error {

	if p.Title == repository.DuplicateKey {
		return errors.New("duplicate key value")
	}

	if p.Title == repository.UnexpectedKey {
		return errors.New("some error")
	}

	if p.Title == repository.NotFoundKey {
		return sql.ErrNoRows
	}

	return nil
}

func (r *mockPostRepo) DeletePost(id int) error {
	if id == repository.UnexpectedKeyInt {
		return errors.New("some error")
	}

	return nil
}

func (r *mockPostRepo) CountPosts(filters models.PostsFilters) (int, error) {
	if filters.Order == repository.UnexpectedKey {
		return 0, errors.New("some error")
	}

	return 1, nil
}

func (r *mockPostRepo) GetCategories() ([]models.Category, error) {
	return nil, nil
}

func (r *mockPostRepo) GetCategoryById(id int) (models.Category, error) {
	var category models.Category

	if id == repository.NotFoundKeyInt {
		return category, sql.ErrNoRows
	}

	if id == repository.UnexpectedKeyInt {
		return category, errors.New("some error")
	}

	return category, nil
}

func (r *mockPostRepo) GetCategoryBySlug(slug string) (models.Category, error) {
	var category models.Category

	if slug == repository.NotFoundKey {
		return category, sql.ErrNoRows
	}

	if slug == repository.UnexpectedKey {
		return category, errors.New("some error")
	}

	return category, nil
}
