package repo

import (
	"context"
	"goods/internal/model"

	"github.com/jmoiron/sqlx"
)

type CustomersRepo struct {
	db *sqlx.DB
}

func NewCustomersRepo(db *sqlx.DB) *CustomersRepo {
	return &CustomersRepo{db: db}
}

func (r *CustomersRepo) CreateCustomer(ctx context.Context, customer model.Customer) (int64, error) {
	var customerID int64
	query := "INSERT INTO customers(user_id, customer_full_name, customer_phone) VALUES(:user_id, :customer_full_name, :customer_phone) RETURNING customer_id"
	if err := r.db.QueryRowxContext(ctx, query, &customer).Scan(&customerID); err != nil {
		return -1, err
	}
	return customerID, nil
}

func (r *CustomersRepo) GetCustomerByID(ctx context.Context, customerID int64) (model.Customer, error) {
	var customer model.Customer
	query := "SELECT * FROM customers WHERE customer_id = $1"
	if err := r.db.Get(&customer, query, customerID); err != nil {
		return model.Customer{}, err
	}
	return customer, nil
}