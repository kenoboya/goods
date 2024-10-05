package repo

import (
	"context"
	"goods/internal/model"

	"github.com/jmoiron/sqlx"
)

type PromocodesRepo struct {
	db *sqlx.DB
}

func NewPromocodesRepo(db *sqlx.DB) *PromocodesRepo {
	return &PromocodesRepo{db: db}
}

func (r *PromocodesRepo) GetPromocodeByID(ctx context.Context, promocodeID string) (model.Promocode, error) {
	var promocode model.Promocode
	query := "SELECT * FROM promocodes WHERE promocode_id = $1"
	if err := r.db.GetContext(ctx, &promocode, query, promocodeID); err != nil {
		return model.Promocode{}, err
	}
	return promocode, nil
}
