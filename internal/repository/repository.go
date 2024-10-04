package repo

import (
	"context"
	"goods/internal/model"
)

type Repositories struct {
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
	GetProductByID(ctx context.Context, productID int) (model.Product, error)
}

type Baskets interface {
	CreateBasket(ctx context.Context, customerID string, product_id int) error
	UpdateProductFromBasket(ctx context.Context, customerID string, productID int, quantity int8) error
	DeleteProductFromBasket(ctx context.Context, customerID string, productID int) error
}

type Orders interface {
	CreateOrder(ctx context.Context, order model.OrderDatabase) (int64, error)
	CreateOrderDetails(ctx context.Context, order model.OrderDetailsDatabase) error
	CreateOrderProducts(ctx context.Context, order model.OrderProductDatabase) error
	// GetOrdersDetailsByCustomerID(ctx context.Context, customerID string) ([]model.OrderDetails, error)
}
