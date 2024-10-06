package service

import (
	"context"
	"goods/internal/model"
	repo "goods/internal/repository"
)

type ProductsService struct {
	productsRepo repo.Products
}

func NewProductsService(productsRepo repo.Products) *ProductsService {
	return &ProductsService{productsRepo: productsRepo}
}

func (s *ProductsService) GetProducts(ctx context.Context) ([]model.Product, error) {
	return s.productsRepo.GetProducts(ctx)
}

func (s *ProductsService) GetProductsByCategoryID(ctx context.Context, categoryID int8) ([]model.Product, error) {
	return s.productsRepo.GetProductsByCategoryID(ctx, categoryID)
}

func (s *ProductsService) GetProductsByID(ctx context.Context, productID int) (model.Product, error) {
	return s.productsRepo.GetProductByID(ctx, productID)
}
func (s *ProductsService) GetProductByName(ctx context.Context, productName string) (model.Product, error) {
	return s.productsRepo.GetProductByName(ctx, productName)
}
