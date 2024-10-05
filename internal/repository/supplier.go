package repo

import (
	"context"
	"goods/internal/model"

	"github.com/jmoiron/sqlx"
)

type SuppliersRepo struct {
	db *sqlx.DB
}

func NewSuppliersRepo(db *sqlx.DB) *SuppliersRepo {
	return &SuppliersRepo{db: db}
}

func (r *SuppliersRepo) CreateSupplier(ctx context.Context, supplier model.Supplier) error {
	query := "INSERT INTO suppliers(company_name, contact_name, contact_title, address, phone, contract) VALUES(:company_name, :contact_name, :contact_title, :address, :phone, :contract)"
	_, err := r.db.NamedExecContext(ctx, query, &supplier)
	if err != nil {
		return err
	}
	return nil
}
