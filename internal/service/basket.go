package service

import (
	"context"
	"goods/internal/model"
	repo "goods/internal/repository"
)

type BasketsService struct {
	basketsRepo repo.Baskets
}

func NewBasketsService(basketsRepo repo.Baskets) *BasketsService {
	return &BasketsService{basketsRepo: basketsRepo}
}

func (s *BasketsService) AddProduct(ctx context.Context, customerID string, productRequest model.ProductRequest) error {
	return s.basketsRepo.CreateBasket(ctx, customerID, productRequest)
}

func (s *BasketsService) GetProducts(ctx context.Context, customerID string) ([]model.Product, error) {
	return s.basketsRepo.GetProductsFromBasket(ctx, customerID)
}

func (s *BasketsService) UpdateProduct(ctx context.Context, customerID string, productRequest model.ProductRequest) error {
	return s.basketsRepo.UpdateProductFromBasket(ctx, customerID, productRequest)
}

func (s *BasketsService) DeleteProduct(ctx context.Context, customerID string, productID int) error {
	return s.basketsRepo.DeleteProductFromBasket(ctx, customerID, productID)
}
