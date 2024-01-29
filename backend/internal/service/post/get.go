package post

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/Noblefel/ManorTalk/backend/internal/models"
)

func (s *postService) Get(slug string) (models.Post, error) {
	post, err := s.postRepo.GetPostBySlug(slug)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return post, ErrNoPost
		}

		return post, fmt.Errorf("%s: %w", "Error getting post by slug", err)
	}

	return post, nil
}

func (s *mockPostService) Get(slug string) (models.Post, error) {
	var post models.Post

	if slug == ErrNoPost.Error() {
		return post, ErrNoPost
	}

	if slug == http.StatusText(http.StatusInternalServerError) {
		return post, errors.New("unexpected error")
	}

	return post, nil
}
