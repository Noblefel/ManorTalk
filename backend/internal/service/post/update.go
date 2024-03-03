package post

import (
	"database/sql"
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"time"

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

	post = models.Post{
		Id:         post.Id,
		Title:      payload.Title,
		Slug:       slug.Make(payload.Title),
		Excerpt:    payload.Excerpt,
		Content:    payload.Content,
		CategoryId: payload.CategoryId,
	}

	files, ok := payload.Files["image"]
	if ok {
		f, err := files[0].Open()
		if err != nil {
			return fmt.Errorf("opening file: %w", err)
		}
		defer f.Close()

		post.Image = fmt.Sprintf(
			"%d%d-%s%s",
			time.Now().UnixNano(),
			authId,
			uuid.New(),
			filepath.Ext(files[0].Filename),
		)

		err = img.Upload(f, "images/post/"+post.Image)
		if err != nil {
			switch err {
			case img.ErrTooLarge:
				return ErrImageTooLarge
			case img.ErrType:
				return ErrImageInvalid
			default:
				return fmt.Errorf("uploading image: %w", err)
			}
		}
	}

	err = s.postRepo.UpdatePost(post)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return ErrDuplicateTitle
		}

		return fmt.Errorf("updating post: %w", err)
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
