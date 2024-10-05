package repo

import (
	"context"
	"goods/internal/model"

	"github.com/jmoiron/sqlx"
)

type ProductsRepo struct {
	db *sqlx.DB
}

func NewProductsRepo(db *sqlx.DB) *ProductsRepo {
	return &ProductsRepo{db: db}
}

func (r *ProductsRepo) GetProducts(ctx context.Context) ([]model.Product, error) {
	var products []model.Product
	query := "SELECT * FROM products"
	if err := r.db.Select(&products, query); err != nil {
		return []model.Product{}, err
	}
	return products, nil
}

func (r *ProductsRepo) GetProductsByCategoryID(ctx context.Context, categoryID int8) ([]model.Product, error) {
	var products []model.Product
	query := "SELECT * FROM products WHERE category_id = $1"
	if err := r.db.Select(&products, query, categoryID); err != nil {
		return []model.Product{}, err
	}
	return products, nil
}

func (r *ProductsRepo) GetProductByID(ctx context.Context, productID int) (model.Product, error) {
	var product model.Product
	query := "SELECT * FROM products WHERE product_id = $1"
	if err := r.db.Get(&product, query, productID); err != nil {
		return model.Product{}, err
	}
	return product, nil
}
