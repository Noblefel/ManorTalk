package post

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Noblefel/ManorTalk/backend/internal/models"
	"github.com/Noblefel/ManorTalk/backend/internal/utils/img"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

func (s *postService) Update(payload models.PostUpdateInput, urlSlug string, authId int) error {
	post, err := s.postRepo.GetPostBySlug(urlSlug)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return ErrNoPost
		}

		return fmt.Errorf("getting post by slug: %w", err)
	}

	if authId != post.UserId {
		return ErrUnauthorized
	}

	if payload.CategoryId != post.CategoryId {
		_, err = s.postRepo.GetCategoryById(payload.CategoryId)
		if err != nil {
			if errors.Is(sql.ErrNoRows, err) {
				return ErrNoCategory
			}

			return fmt.Errorf("getting category by id: %w", err)
		}
	}

	var oldImage string

	if payload.Image != nil {
		ext, err := img.Verify(payload.Image)
		if err != nil {
			switch err {
			case img.ErrTooLarge:
				return ErrImageTooLarge
			case img.ErrType:
				return ErrImageInvalid
			default:
				return fmt.Errorf("verifying image: %w", err)
			}
		}

		name := fmt.Sprintf("%s-%d", uuid.New(), authId) + ext
		oldImage, post.Image = post.Image, name

		err = img.Save(payload.Image, filepath.Join("images", "post", post.Image))
		if err != nil {
			return fmt.Errorf("saving image: %w", err)
		}
	}

	post.Title = payload.Title
	post.Slug = slug.Make(payload.Title)
	post.Excerpt = payload.Excerpt
	post.Content = payload.Content
	post.CategoryId = payload.CategoryId

	if err := s.postRepo.UpdatePost(post); err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return ErrDuplicateTitle
		}

		return fmt.Errorf("updating post: %w", err)
	}

	if oldImage != "" {
		if err := os.Remove(filepath.Join("images", "post", oldImage)); err != nil {
			log.Println("unable to delete image: ", err)
		}
	}

	return nil
}

func (s *mockPostService) Update(payload models.PostUpdateInput, urlSlug string, authId int) error {
	switch urlSlug {
	case ErrNoPost.Error():
		return ErrNoPost
	case ErrNoCategory.Error():
		return ErrNoCategory
	case ErrUnauthorized.Error():
		return ErrUnauthorized
	case ErrDuplicateTitle.Error():
		return ErrDuplicateTitle
	case ErrImageTooLarge.Error():
		return ErrImageTooLarge
	case ErrImageInvalid.Error():
		return ErrImageInvalid
	case "unexpected error":
		return errors.New("unexpected error")
	default:
		return nil
	}
}
