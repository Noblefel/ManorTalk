package post

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Noblefel/ManorTalk/backend/internal/models"
	"github.com/gosimple/slug"
)

func (s *postService) Create(payload models.PostCreateInput) (models.Post, error) {
	var post models.Post

	category, err := s.postRepo.GetCategoryById(payload.CategoryId)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return post, ErrNoCategory
		}

		return post, fmt.Errorf("%s: %w", "Error getting category by id", err)
	}

	post = models.Post{
		UserId:     1,
		Title:      payload.Title,
		Slug:       slug.Make(payload.Title),
		Excerpt:    payload.Excerpt,
		Content:    payload.Content,
		CategoryId: payload.CategoryId,
	}

	post, err = s.postRepo.CreatePost(post)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return post, ErrDuplicateTitle
		}

		return post, fmt.Errorf("%s: %w", "Error creating post", err)
	}

	post.Category = category

	return post, nil
}

func (s *mockPostService) Create(payload models.PostCreateInput) (models.Post, error) {
	var post models.Post

	if payload.Title == ErrNoCategory.Error() {
		return post, ErrNoCategory
	}

	if payload.Title == ErrDuplicateTitle.Error() {
		return post, ErrDuplicateTitle
	}

	if payload.Title == http.StatusText(http.StatusInternalServerError) {
		return post, errors.New("unexpected error")
	}

	return post, nil
}
