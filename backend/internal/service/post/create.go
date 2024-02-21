package post

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/Noblefel/ManorTalk/backend/internal/models"
	"github.com/gosimple/slug"
)

func (s *postService) Create(payload models.PostCreateInput, userId int) (models.Post, error) {
	var post models.Post

	category, err := s.postRepo.GetCategoryById(payload.CategoryId)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return post, ErrNoCategory
		}

		return post, fmt.Errorf("%s: %w", "Error getting category by id", err)
	}

	post = models.Post{
		UserId:     userId,
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

func (s *mockPostService) Create(payload models.PostCreateInput, userId int) (models.Post, error) {
	var post models.Post
	switch payload.Title {
	case ErrNoCategory.Error():
		return post, ErrNoCategory
	case ErrDuplicateTitle.Error():
		return post, ErrDuplicateTitle
	case "unexpected error":
		return post, errors.New("unexpected error")
	default:
		return post, nil
	}
}
