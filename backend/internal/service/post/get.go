package post

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Noblefel/ManorTalk/backend/internal/models"
)

func (s *postService) Get(slug string) (models.Post, error) {
	post, err := s.postRepo.GetPostBySlug(slug)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return post, ErrNoPost
		}

		return post, fmt.Errorf("getting post by slug: %w", err)
	}

	return post, nil
}

func (s *mockPostService) Get(slug string) (models.Post, error) {
	var post models.Post
	switch slug {
	case ErrNoPost.Error():
		return post, ErrNoPost
	case "unexpected error":
		return post, errors.New("unexpected error")
	default:
		return post, nil
	}
}
