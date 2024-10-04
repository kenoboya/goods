package repo

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type BasketsRepo struct {
	db *sqlx.DB
}

func NewBasketsRepo(db *sqlx.DB) *BasketsRepo {
	return &BasketsRepo{db: db}
}

func (r *BasketsRepo) CreateBasket(ctx context.Context, customerID string, product_id int) error {
	query := "INSERT INTO baskets (customer_id, product_id) VALUES ($1, $2)"
	if _, err := r.db.ExecContext(ctx, query, customerID, product_id); err != nil {
		return err
	}
	return nil
}

func (r *BasketsRepo) UpdateProductFromBasket(ctx context.Context, customerID string, productID int, quantity int8) error {
	query := "UPDATE baskets SET quantity = $1 WHERE customer_id = $2 AND product_id = $3"
	if _, err := r.db.ExecContext(ctx, query, quantity, customerID, productID); err != nil {
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
