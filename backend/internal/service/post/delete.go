package post

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
)

func (s *postService) Delete(slug string) error {
	post, err := s.postRepo.GetPostBySlug(slug)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return ErrNoPost
		}

		return fmt.Errorf("%s: %w", "Error getting post by slug", err)
	}

	err = s.postRepo.DeletePost(post.Id)
	if err != nil {
		return fmt.Errorf("%s: %w", "Error deleting post", err)
	}

	return nil
}

func (s *mockPostService) Delete(slug string) error {
	if slug == ErrNoPost.Error() {
		return ErrNoPost
	}

	if slug == http.StatusText(http.StatusInternalServerError) {
		return errors.New("unexpected error")
	}

	return nil
}
