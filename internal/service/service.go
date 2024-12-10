package service

import (
	"context"
	"goods/internal/model"
	repo "goods/internal/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Services struct {
	Categories Categories
	Products   Products
	Baskets    Baskets
	Orders     Orders
	Customers  Customers
	// Suppliers  Suppliers
	// Shipping   Shipping
	// Promocodes Promocodes
}

func NewServices(repositories *repo.Repositories) *Services {
	return &Services{
		Categories: NewCategoriesService(repositories.Categories),
		Products:   NewProductsService(repositories.Products),
		Baskets:    NewBasketsService(repositories.Baskets),
		Orders: NewOrdersService(
			repositories.Orders,
			repositories.Customers,
			repositories.Shipping),
		Customers: repositories.Customers,
	}
}

type Categories interface {
	CreateCategories(ctx context.Context, category model.CreateCategoryRequest) error
	GetCategories(ctx context.Context) ([]model.Category, error)
	GetCategoryByID(ctx context.Context, categoryID int8) (model.Category, error)
	GetCategoryByName(ctx context.Context, categoryName string) (model.Category, error)
}

type Products interface {
	CreateProduct(ctx context.Context, product model.CreateProductRequest) error
	GetProducts(ctx context.Context) ([]model.Product, error)
	GetProductsByCategoryID(ctx context.Context, categoryID int8) ([]model.Product, error)
	GetProductsByID(ctx context.Context, productID int) (model.Product, error)
	GetProductByName(ctx context.Context, productName string) (model.Product, error)
}

type Baskets interface {
	GetProducts(ctx context.Context, customerID string) ([]model.Product, error)
	AddProduct(ctx context.Context, customerID string, productRequest model.ProductRequest) error
	UpdateProduct(ctx context.Context, customerID string, productRequest model.ProductRequest) error
	DeleteProduct(ctx context.Context, customerID string, productID int) error
}

type Orders interface {
	CreateOrder(ctx context.Context, order model.OrderRequest) (int64, error)
	GetTotalSumOrder(ctx context.Context, orderID int64) (float64, error)
	LinkTransactionToOrder(ctx context.Context, orderID int64, transactionID string) error
}

type Customers interface {
	GetCustomers(ctx context.Context) ([]model.Customer, error)
	GetCustomerByID(ctx context.Context, customerID int64) (model.Customer, error)
}

type Suppliers interface {
	CreateSupplier(ctx context.Context, supplier model.Supplier) error
}

type Shipping interface {
	CreateShippingDetails(ctx context.Context, shippingDetails model.DeliveryAddress) (int64, error)
}

type Promocodes interface {
	GetPromocodeByID(ctx context.Context, promocodeID string) (model.Promocode, error)
}
