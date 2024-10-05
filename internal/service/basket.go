package service

import (
	"context"
	repo "goods/internal/repository"
)

type BasketsService struct {
	basketsRepo repo.Baskets
}

func NewBasketsService(basketsRepo repo.Baskets) *BasketsService {
	return &BasketsService{basketsRepo: basketsRepo}
}

func (s *BasketsService) AddProduct(ctx context.Context, customerID string, productID int) error {
	return s.basketsRepo.CreateBasket(ctx, customerID, productID)
}

func (s *BasketsService) UpdateProduct(ctx context.Context, customerID string, productID int, quantity int8) error {
	return s.basketsRepo.UpdateProductFromBasket(ctx, customerID, productID, quantity)
}

func (s *BasketsService) DeleteProduct(ctx context.Context, customerID string, productID int) error {
	return s.basketsRepo.DeleteProductFromBasket(ctx, customerID, productID)
}
