package model

import "time"

type DeliveryAddress struct {
	ShipRegion  string  `json:"ship_region" db:"ship_region"`
	ShipCity    string  `json:"ship_city" db:"ship_city"`
	ShipAddress string  `json:"ship_address" db:"ship_address"`
	Porch       *int8   `json:"porch" db:"porch"`
	Floor       *int8   `json:"floor" db:"floor"`
	Apartment   *int8   `json:"apartment" db:"apartment"`
	Intercom    *string `json:"intercom" db:"intercom"`
	Description *string `json:"description" db:"description"`
}

type Customer struct {
	UserID           *string `json:"user_id" db:"user_id"`
	CustomerFullName string  `json:"customer_full_name" db:"customer_full_name"`
	CustomerPhone    string  `json:"customer_phone" db:"customer_phone"`
}

type OrderBriefInfo struct {
	OrderID int64 `json:"order_id" db:"order_id"`
}

type OrderDetails struct {
	OrderBriefInfo
	Products        []ProductResponse `json:"Products"`
	Customer        Customer          `json:"Customer"`
	DeliveryAddress DeliveryAddress   `json:"DeliveryAddress"`
	Promocode       *Promocode        `json:"Promocode"`
	OrderDate       time.Time         `json:"order_date" db:"order_date"`
}

type OrderRequest struct {
	Customer        Customer
	TransactionID   string
	Products        []ProductRequest
	DeliveryAddress DeliveryAddress
	Promocode       *string
}

type OrderDatabase struct {
	TransactionID string
	CustomerID    int
}

type OrderDetailsDatabase struct {
	OrderID           int64
	ShippingDetailsID int64
	PromocodeID       *string
}

type OrderProductDatabase struct {
	OrderID   int64
	ProductID int
	Quantity  int8
}
