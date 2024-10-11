package repo

import (
	"context"
	"database/sql"
	"goods/internal/model"

	"github.com/jmoiron/sqlx"
)

type ProductsRepo struct {
	db *sqlx.DB
}

func NewProductsRepo(db *sqlx.DB) *ProductsRepo {
	return &ProductsRepo{db: db}
}

func (r *ProductsRepo) CreateProduct(ctx context.Context, product model.CreateProductRequest) error {
	query := "INSERT INTO products(product_name, supplier_id, category_id, unit_price, stock, discount, quantity_per_unit, weight, image) VALUES(:product_name, :supplier_id, :category_id, :unit_price, :stock, :discount, :quantity_per_unit, :weight, :image)"
	_, err := r.db.NamedExecContext(ctx, query, &product)
	if err != nil {
		return err
	}
	return nil
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
		if err == sql.ErrNoRows {
			return []model.Product{}, model.ErrNotFoundProducts
		}
		return []model.Product{}, err
	}
	return products, nil
}

func (r *ProductsRepo) GetProductsByCategory(ctx context.Context, categoryID int8) ([]model.Product, error) {
	var products []model.Product
	query := "SELECT * FROM products WHERE category_id = $1"
	if err := r.db.Select(&products, query, categoryID); err != nil {
		if err == sql.ErrNoRows {
			return []model.Product{}, model.ErrNotFoundProducts
		}
		return []model.Product{}, err
	}
	return products, nil
}

func (r *ProductsRepo) GetProductByID(ctx context.Context, productID int) (model.Product, error) {
	var product model.Product
	query := "SELECT * FROM products WHERE product_id = $1"
	if err := r.db.Get(&product, query, productID); err != nil {
		if err == sql.ErrNoRows {
			return model.Product{}, model.ErrNotFoundProduct
		}
		return model.Product{}, err
	}
	return product, nil
}

func (r *ProductsRepo) GetProductByName(ctx context.Context, productName string) (model.Product, error) {
	var product model.Product
	query := "SELECT * FROM products WHERE product_name = $1"
	if err := r.db.Get(&product, query, productName); err != nil {
		if err == sql.ErrNoRows {
			return model.Product{}, model.ErrNotFoundProduct
		}
		return model.Product{}, err
	}
	return product, nil
}
