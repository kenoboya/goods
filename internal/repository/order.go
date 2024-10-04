package repo

import (
	"context"
	"goods/internal/model"

	"github.com/jmoiron/sqlx"
)

type OrdersRepo struct {
	db *sqlx.DB
}

func NewOrdersRepo(db *sqlx.DB) *OrdersRepo {
	return &OrdersRepo{db: db}
}

func (r *OrdersRepo) CreateOrder(ctx context.Context, order model.OrderDatabase) (int64, error) {
	var orderID int64
	query := "INSERT INTO orders (customer_id, transaction_id) VALUES (:customer_id, :transaction_id) RETURNING order_id"
	if err := r.db.QueryRowxContext(ctx, query, &order).Scan(&orderID); err != nil {
		return -1, err
	}
	return orderID, nil
}

func (r *OrdersRepo) CreateOrderDetails(ctx context.Context, order model.OrderDetailsDatabase) error {
	query := "INSERT INTO orders_details(order_id, shipping_details_id, promocode_id) VALUES(:order_id, :shipping_details_id, :promocode_id)"
	if _, err := r.db.NamedExecContext(ctx, query, &order); err != nil {
		return err
	}
	return nil
}

func (r *OrdersRepo) CreateOrderProducts(ctx context.Context, order model.OrderProductDatabase) error {
	query := "INSERT INTO orders_products(order_id, product_id, quantity) VALUES(:order_id, :product_id, :quantity)"
	if _, err := r.db.NamedExecContext(ctx, query, &order); err != nil {
		return err
	}
	return nil
}
