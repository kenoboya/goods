package repo

import (
	"context"
	"goods/internal/model"

	"github.com/jmoiron/sqlx"
)

type BasketsRepo struct {
	db *sqlx.DB
}

func NewBasketsRepo(db *sqlx.DB) *BasketsRepo {
	return &BasketsRepo{db: db}
}

func (r *BasketsRepo) CreateBasket(ctx context.Context, customerID string, productRequest model.ProductRequest) error {
	query := "INSERT INTO baskets (customer_id, product_id, quantity) VALUES ($1, $2, $3)"
	if _, err := r.db.ExecContext(ctx, query, customerID, productRequest.ProductID, productRequest.Quantity); err != nil {
		return err
	}
	return nil
}

func (r *BasketsRepo) GetProductsFromBasket(ctx context.Context, customerID string) ([]model.Product, error) {
	var products []model.Product
	query := `
	SELECT 
		p.product_id, p.product_name, p.supplier_id, 
		p.category_id, p.unit_price, p.in_stock, p.discount, 
		p.quantity_per_unit, p.weight, p.image
	FROM 
		baskets b
	JOIN 
		products p ON b.product_id = p.product_id
	WHERE 
		b.customer_id = $1
`
	if err := r.db.Select(&products, query, customerID); err != nil {
		return []model.Product{}, err
	}
	return products, nil
}

func (r *BasketsRepo) UpdateProductFromBasket(ctx context.Context, customerID string, productRequest model.ProductRequest) error {
	query := "UPDATE baskets SET quantity = $1 WHERE customer_id = $2 AND product_id = $3"
	if _, err := r.db.ExecContext(ctx, query, productRequest.Quantity, customerID, productRequest.ProductID); err != nil {
		return err
	}
	return nil
}

func (r *BasketsRepo) DeleteProductFromBasket(ctx context.Context, customerID string, productID int) error {
	query := "DELETE FROM baskets WHERE customer_id = $1 AND product_id = $2"
	if _, err := r.db.ExecContext(ctx, query, customerID, productID); err != nil {
		return err
	}
	return nil
}
