package repo

import (
	"context"
	"goods/internal/model"

	"github.com/jmoiron/sqlx"
)

type ShippingRepo struct {
	db *sqlx.DB
}

func NewShippingRepo(db *sqlx.DB) *ShippingRepo {
	return &ShippingRepo{db: db}
}

func (r *ShippingRepo) CreateShippingDetails(ctx context.Context, shippingDetails model.DeliveryAddress) (int64, error) {
	var shippingDetailsID int64
	query := `INSERT INTO shipping_details(ship_region, ship_city, ship_address, porch, floor, apartment, intercom, description)
              VALUES(:ship_region, :ship_city, :ship_address, :porch, :floor, :apartment, :intercom, :description)
              RETURNING shipping_details_id`

	if err := r.db.QueryRowxContext(ctx, query, &shippingDetails).Scan(&shippingDetailsID); err != nil {
		return -1, err
	}
	return shippingDetailsID, nil
}
