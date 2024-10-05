package service

import (
	"context"
	"goods/internal/model"
	repo "goods/internal/repository"
)

type CategoriesService struct {
	categoriesRepo repo.Categories
}

func NewCategoriesService(categoriesRepo repo.Categories) *CategoriesService {
	return &CategoriesService{categoriesRepo: categoriesRepo}
}

func (s *CategoriesService) GetCategories(ctx context.Context) ([]model.Category, error) {
	return s.categoriesRepo.GetCategories(ctx)
}

func (s *CategoriesService) GetCategoryByID(ctx context.Context, categoryID int8) (model.Category, error) {
	return s.categoriesRepo.GetCategoryByID(ctx, categoryID)
}
