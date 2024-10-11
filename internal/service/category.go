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

func (s *CategoriesService) CreateCategories(ctx context.Context, category model.CreateCategoryRequest) error {
	return s.categoriesRepo.CreateCategories(ctx, category)
}

func (s *CategoriesService) GetCategories(ctx context.Context) ([]model.Category, error) {
	return s.categoriesRepo.GetCategories(ctx)
}

func (s *CategoriesService) GetCategoryByID(ctx context.Context, categoryID int8) (model.Category, error) {
	return s.categoriesRepo.GetCategoryByID(ctx, categoryID)
}

func (s *CategoriesService) GetCategoryByName(ctx context.Context, categoryName string) (model.Category, error) {
	return s.categoriesRepo.GetCategoryByName(ctx, categoryName)
}
