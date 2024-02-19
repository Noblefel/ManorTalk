package post

import (
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"strconv"

	"github.com/Noblefel/ManorTalk/backend/internal/models"
	"github.com/Noblefel/ManorTalk/backend/internal/utils/pagination"
	"github.com/gosimple/slug"
)

func (s *postService) GetMany(q url.Values) ([]models.Post, *pagination.Meta, error) {
	var posts []models.Post
	var pgMeta *pagination.Meta
	var err error

	cursor, _ := strconv.Atoi(q.Get("cursor"))
	uId, _ := strconv.Atoi(q.Get("user"))
	limit, err := strconv.Atoi(q.Get("limit"))
	if err != nil {
		limit = 10
	}

	filters := models.PostsFilters{
		Order:    q.Get("order"),
		Category: q.Get("category"),
		Search:   q.Get("search"),
		Cursor:   cursor,
		UserId:   uId,
		Limit:    limit,
	}

	if filters.Category != "" {
		c, err := s.postRepo.GetCategoryBySlug(q.Get("category"))
		if err != nil {
			if errors.Is(sql.ErrNoRows, err) {
				return posts, nil, ErrNoCategory
			}

			return posts, nil, fmt.Errorf("getting category by slug: %w", err)
		}

		filters.Category = c.Slug
	}

	if cursor == 0 {
		pgMeta, err = pagination.NewMeta(q, limit)
		if err != nil {
			return posts, pgMeta, fmt.Errorf("creating pagination meta: %w", err)
		}

		// Client could optionally attach "total" parameter to the url.
		// This skip the below statement to reduce further bottleneck
		if pgMeta.Total == 0 {
			total, err := s.postRepo.CountPosts(filters)
			if err != nil && !errors.Is(sql.ErrNoRows, err) {
				return posts, nil, fmt.Errorf("counting posts: %w", err)
			}

			pgMeta.SetNewTotal(total, limit)
		}
	}

	posts, err = s.postRepo.GetPosts(pgMeta, filters)
	if err != nil && !errors.Is(sql.ErrNoRows, err) {
		return posts, nil, fmt.Errorf("getting posts: %w", err)
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
