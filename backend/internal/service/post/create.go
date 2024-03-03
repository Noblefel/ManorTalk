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

func (s *postService) Create(payload models.PostCreateInput, authId int) (models.Post, error) {
	var post models.Post

	category, err := s.postRepo.GetCategoryById(payload.CategoryId)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return post, ErrNoCategory
		}

		return post, fmt.Errorf("getting category by id: %w", err)
	}

	post = models.Post{
		UserId:     authId,
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
			return post, fmt.Errorf("opening file: %w", err)
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
				return post, ErrImageTooLarge
			case img.ErrType:
				return post, ErrImageInvalid
			default:
				return post, fmt.Errorf("uploading image: %w", err)
			}
		}
	}

	post, err = s.postRepo.CreatePost(post)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return post, ErrDuplicateTitle
		}

		return post, fmt.Errorf("creating post: %w", err)
	}

	post.Category = category

	return post, nil
}

func (s *mockPostService) Create(payload models.PostCreateInput, authId int) (models.Post, error) {
	var post models.Post
	switch payload.Title {
	case ErrNoCategory.Error():
		return post, ErrNoCategory
	case ErrDuplicateTitle.Error():
		return post, ErrDuplicateTitle
	case ErrImageTooLarge.Error():
		return post, ErrImageTooLarge
	case ErrImageInvalid.Error():
		return post, ErrImageInvalid
	case "unexpected error":
		return post, errors.New("unexpected error")
	default:
		return post, nil
	}
}
