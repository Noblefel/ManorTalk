package post

import (
	"database/sql"
	"errors"
	"fmt"
	"net/url"

	"github.com/Noblefel/ManorTalk/backend/internal/models"
	"github.com/Noblefel/ManorTalk/backend/internal/utils/pagination"
	"github.com/gosimple/slug"
)

func (s *postService) GetMany(q url.Values) ([]models.Post, *pagination.Meta, error) {
	var posts []models.Post

	filters := models.PostsFilters{
		Order:    q.Get("order"),
		Category: q.Get("category"),
		Search:   q.Get("search"),
	}

	if filters.Category != "" {
		c, err := s.postRepo.GetCategoryBySlug(q.Get("category"))
		if err != nil {
			if errors.Is(sql.ErrNoRows, err) {
				return posts, nil, ErrNoCategory
			}

			return posts, nil, fmt.Errorf("%s: %w", "Error getting category by slug", err)
		}

		filters.Category = c.Slug
	}

	pgMeta, err := pagination.NewMeta(q)
	if err != nil {
		return posts, pgMeta, fmt.Errorf("%s: %w", "Error creating pagination meta", err)
	}

	// After the initial query and wants to navigate to the next page,
	// client should attach the "total" parameter to the url.
	// This will skip the below statement to reduce further bottleneck
	// Or
	// if the requested posts doesn't need pagination, for example to get the
	// latest posts to be displayed on the sidebar. In this case, feel free
	// to put any number on the "total" param as long as it's not 0.
	if pgMeta.Total == 0 {
		total, err := s.postRepo.CountPosts(filters)
		if err != nil && !errors.Is(sql.ErrNoRows, err) {
			return posts, nil, fmt.Errorf("%s: %w", "Error counting posts", err)
		}

		pgMeta.SetNewTotal(total)
	}

	posts, err = s.postRepo.GetPosts(pgMeta, filters)
	if err != nil && !errors.Is(sql.ErrNoRows, err) {
		return posts, nil, fmt.Errorf("%s: %w", "Error getting posts", err)
	}

	return posts, pgMeta, nil
}

func (s *mockPostService) GetMany(q url.Values) ([]models.Post, *pagination.Meta, error) {
	posts, pgMeta := []models.Post{}, &pagination.Meta{}

	if q.Has(slug.Make(ErrNoCategory.Error())) {
		return posts, nil, ErrNoCategory
	}

	if q.Has(slug.Make("unexpected error")) {
		return posts, nil, errors.New("unexpected error")
	}

	return posts, pgMeta, nil
}
