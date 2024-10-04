package service

import (
	"context"
	"goods/internal/model"
)

type Services struct {
	Categories Categories
	Products   Products
	Baskets    Baskets
	Orders     Orders
}

type Categories interface {
	GetCategories(ctx context.Context) ([]model.Category, error)
	GetCategoryByID(ctx context.Context, categoryID int8) (model.Category, error)
}

type Products interface {
	GetProducts(ctx context.Context) ([]model.Product, error)
	GetProductsByCategoryID(ctx context.Context, categoryID int8) ([]model.Product, error)
	GetProductsByID(ctx context.Context, productID int) (model.Product, error)
}

type Baskets interface {
	AddProduct(ctx context.Context, customerID string, product_id int) error
	UpdateProduct(ctx context.Context, customerID string, product_id int, quantity int8) error
	DeleteProduct(ctx context.Context, customerID string, product_id int) error
}

type Orders interface {
	SaveOrder(ctx context.Context, order model.OrderRequest) error
}
