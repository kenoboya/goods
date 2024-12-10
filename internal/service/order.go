package service

import (
	"context"
	"goods/internal/model"
	repo "goods/internal/repository"
	logger "goods/pkg/logger/zap"

	"go.uber.org/zap"
)

type OrdersService struct {
	ordersRepo    repo.Orders
	customersRepo repo.Customers
	shippingRepo  repo.Shipping
}

func NewOrdersService(ordersRepo repo.Orders,
	customersRepo repo.Customers, shippingRepo repo.Shipping) *OrdersService {
	return &OrdersService{
		ordersRepo:    ordersRepo,
		customersRepo: customersRepo,
		shippingRepo:  shippingRepo,
	}
}

func (s *OrdersService) CreateOrder(ctx context.Context, order model.OrderRequest) (int64, error) {
	customerID, err := s.customersRepo.CreateCustomer(ctx, order.Customer)
	if err != nil {
		logger.Error("Failed to create customer in the repository",
			zap.String("action", "CreateCustomer()"),
			zap.Stringp("user_id", order.Customer.UserID),
			zap.String("customer_full_name", order.Customer.CustomerFullName),
			zap.String("customer_phone", order.Customer.CustomerPhone),
			zap.Any("customer_object", order.Customer),
			zap.Error(err),
		)
		return -1, err
	}

	orderID, err := s.ordersRepo.CreateOrder(ctx, model.OrderDatabase{
		CustomerID: customerID,
	})
	if err != nil {
		logger.Error("Failed to create order in the repository",
			zap.String("action", "CreateOrder()"),
			zap.Int64("customer_id", customerID),
			zap.Error(err),
		)
		return -1, err
	}

	shippingDetailsID, err := s.shippingRepo.CreateShippingDetails(ctx, order.DeliveryAddress)
	if err != nil {
		logger.Error("Failed to create shipping details of order in the repository",
			zap.String("action", "CreateShippingDetails()"),
			zap.String("ship_region", order.DeliveryAddress.ShipRegion),
			zap.String("ship_city", order.DeliveryAddress.ShipCity),
			zap.String("ship_address", order.DeliveryAddress.ShipAddress),
			zap.Any("porch", order.DeliveryAddress.Porch),
			zap.Any("floor", order.DeliveryAddress.Floor),
			zap.Any("apartment", order.DeliveryAddress.Apartment),
			zap.Stringp("intercom", order.DeliveryAddress.Intercom),
			zap.Stringp("description", order.DeliveryAddress.Description),
			zap.Error(err),
		)
		return -1, err
	}

	if err := s.ordersRepo.CreateOrderDetails(ctx, model.OrderDetailsDatabase{
		OrderID:           orderID,
		ShippingDetailsID: shippingDetailsID,
		PromocodeID:       order.Promocode,
	}); err != nil {
		logger.Error("Failed to create order details in the repository",
			zap.String("action", "CreateOrderDetails()"),
			zap.Int64("order_id", orderID),
			zap.Int64("shipping_details_id", shippingDetailsID),
			zap.Stringp("promocode", order.Promocode),
			zap.Error(err),
		)
		return -1, err
	}

	for _, productInfo := range order.Products {
		if err := s.ordersRepo.CreateOrderProducts(ctx, model.OrderProductDatabase{
			OrderID:   orderID,
			ProductID: productInfo.ProductID,
			Quantity:  productInfo.Quantity,
		}); err != nil {
			logger.Error("Failed to create relation between order and product in the repository",
				zap.String("action", "CreateOrderProducts()"),
				zap.Int64("order_id", orderID),
				zap.Int("product_id", productInfo.ProductID),
				zap.Any("quantity", productInfo.Quantity),
				zap.Error(err),
			)
			return -1, err
		}
	}
	logger.Infof("Order No. %d has been successfully registered.", orderID)
	return orderID, nil
}

func (s *OrdersService) LinkTransactionToOrder(ctx context.Context, orderID int64, transactionID string) error {
	return s.ordersRepo.UpdateOrderWithTransactionID(ctx, orderID, transactionID)
}

func (s *OrdersService) GetTotalSumOrder(ctx context.Context, orderID int64) (float64, error) {
	return s.ordersRepo.GetTotalSumOrder(ctx, orderID)
}
