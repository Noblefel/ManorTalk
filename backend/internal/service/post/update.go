package post

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/Noblefel/ManorTalk/backend/internal/models"
	"github.com/gosimple/slug"
)

func (s *postService) Update(payload models.PostUpdateInput, urlSlug string) error {
	post, err := s.postRepo.GetPostBySlug(urlSlug)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return ErrNoPost
		}

		return fmt.Errorf("%s: %w", "Error getting post by slug", err)
	}

	post = models.Post{
		Id:         post.Id,
		Title:      payload.Title,
		Slug:       slug.Make(payload.Title),
		Excerpt:    payload.Excerpt,
		Content:    payload.Content,
		CategoryId: payload.CategoryId,
	}

	err = s.postRepo.UpdatePost(post)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return ErrDuplicateTitle
		}

		return fmt.Errorf("%s: %w", "Error updating post", err)
	}

	return nil
}

func (s *mockPostService) Update(payload models.PostUpdateInput, urlSlug string) error {
	switch urlSlug {
	case ErrNoPost.Error():
		return ErrNoPost
	case ErrDuplicateTitle.Error():
		return ErrDuplicateTitle
	case "unexpected error":
		return errors.New("unexpected error")
	default:
		return nil
	}
}
