package repo

import (
	"context"
	"goods/internal/model"

	"github.com/jmoiron/sqlx"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type Repositories struct {
	Categories Categories
	Products   Products
	Baskets    Baskets
	Orders     Orders
	Customers  Customers
	Suppliers  Suppliers
	Shipping   Shipping
	Promocodes Promocodes
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Categories: NewCategoriesRepo(db),
		Products:   NewProductsRepo(db),
		Baskets:    NewBasketsRepo(db),
		Orders:     NewOrdersRepo(db),
		Customers:  NewCustomersRepo(db),
		Suppliers:  NewSuppliersRepo(db),
		Shipping:   NewShippingRepo(db),
		Promocodes: NewPromocodesRepo(db),
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
	GetProductByID(ctx context.Context, productID int) (model.Product, error)
	GetProductByName(ctx context.Context, productName string) (model.Product, error)
}

type Baskets interface {
	CreateBasket(ctx context.Context, customerID string, productRequest model.ProductRequest) error
	GetProductsFromBasket(ctx context.Context, customerID string) ([]model.Product, error)
	UpdateProductFromBasket(ctx context.Context, customerID string, productRequest model.ProductRequest) error
	DeleteProductFromBasket(ctx context.Context, customerID string, productID int) error
}

type Orders interface {
	CreateOrder(ctx context.Context, order model.OrderDatabase) (int64, error)
	CreateOrderDetails(ctx context.Context, order model.OrderDetailsDatabase) error
	CreateOrderProducts(ctx context.Context, order model.OrderProductDatabase) error
	// GetOrdersDetailsByCustomerID(ctx context.Context, customerID string) ([]model.OrderDetails, error)
	UpdateOrderWithTransactionID(ctx context.Context, orderID int64, transactionID string) error
	GetTotalSumOrder(ctx context.Context, orderID int64) (float64, error)
}

type Customers interface {
	CreateCustomer(ctx context.Context, customer model.Customer) (int64, error)
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
