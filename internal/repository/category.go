package repo

import (
	"context"
	"goods/internal/model"

	"github.com/jmoiron/sqlx"
)

type CategoriesRepo struct {
	db *sqlx.DB
}

func NewCategoriesRepo(db *sqlx.DB) *CategoriesRepo {
	return &CategoriesRepo{db: db}
}

func (r *CategoriesRepo) GetCategories(ctx context.Context) ([]model.Category, error) {
	var categories []model.Category
	query := "SELECT * FROM categories"
	if err := r.db.Select(&categories, query); err != nil {
		return []model.Category{}, err
	}
	return categories, nil
}

func (r *CategoriesRepo) GetCategoryByID(ctx context.Context, categoryID int8) (model.Category, error) {
	var category model.Category
	query := "SELECT * FROM categories WHERE category_id = $1"
	if err := r.db.GetContext(ctx, &category, query, categoryID); err != nil {
		return model.Category{}, err
	}
	return category, nil
}

func (r *CategoriesRepo) GetCategoryByName(ctx context.Context, categoryName string) (model.Category, error) {
	var category model.Category
	query := "SELECT * FROM categories WHERE category_name = $1"
	if err := r.db.GetContext(ctx, &category, query, categoryName); err != nil {
		return model.Category{}, err
	}
	return category, nil
}
