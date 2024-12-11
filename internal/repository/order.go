package repo

import (
	"context"
	"database/sql"
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
	query := "INSERT INTO orders (customer_id, payment_method) VALUES (:customer_id, :payment_method) RETURNING order_id"
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

func (r *OrdersRepo) UpdateOrderWithTransactionID(ctx context.Context, orderID int64, transactionID string) error {
	query := "UPDATE ON orders SET transaction_id = $1 WHERE order_id = $2"
	if _, err := r.db.ExecContext(ctx, query, transactionID, orderID); err != nil {
		return err
	}
	return nil
}

func (r *OrdersRepo) GetTotalSumOrder(ctx context.Context, orderID int64) (float64, error) {
	var sum float64
	// procedure plpgsql
	query := `SELECT GetTotalSumOrder($1)`
	if err := r.db.GetContext(ctx, sum, query, orderID); err != nil {
		if err == sql.ErrNoRows {
			return -1, model.ErrNotFoundOrder
		}
		return -1, err
	}
	return sum, nil
}
