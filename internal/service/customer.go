package service

import (
	"context"
	"goods/internal/model"
	repo "goods/internal/repository"
)

type CustomersService struct {
	customersRepo repo.Customers
}

func NewCustomersService(customersRepo repo.Customers) *CustomersService {
	return &CustomersService{customersRepo: customersRepo}
}

func (s *CustomersService) GetCustomers(ctx context.Context) ([]model.Customer, error) {
	return s.customersRepo.GetCustomers(ctx)
}
func (s *CustomersService) GetCustomerByID(ctx context.Context, customerID int64) (model.Customer, error) {
	return s.customersRepo.GetCustomerByID(ctx, customerID)
}
