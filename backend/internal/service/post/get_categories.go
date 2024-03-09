package post

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Noblefel/ManorTalk/backend/internal/models"
)

func (s *postService) GetCategories() ([]models.Category, error) {
	categories, err := s.postRepo.GetCategories()
	if err != nil && !errors.Is(sql.ErrNoRows, err) {
		return categories, fmt.Errorf("getting categories: %w", err)
	}

	return categories, nil
}

func (s *mockPostService) GetCategories() ([]models.Category, error) {
	return []models.Category{}, nil
}
