package post

import (
	"database/sql"
	"errors"
	"fmt"
)

func (s *postService) Delete(slug string, authId int) error {
	post, err := s.postRepo.GetPostBySlug(slug)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return ErrNoPost
		}

		return fmt.Errorf("getting post by slug: %w", err)
	}

	if authId != post.UserId {
		return ErrUnauthorized
	}

	err = s.postRepo.DeletePost(post.Id)
	if err != nil {
		return fmt.Errorf("deleting post: %w", err)
	}

	return nil
}

func (s *mockPostService) Delete(slug string, authId int) error {
	switch slug {
	case ErrNoPost.Error():
		return ErrNoPost
	case ErrUnauthorized.Error():
		return ErrUnauthorized
	case "unexpected error":
		return errors.New("unexpected error")
	default:
		return nil
	}
}
