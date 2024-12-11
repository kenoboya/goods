package model

import "time"

var (
	PaymentMethodCard = "card"
	PaymentMethodCash = "cash"
)

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
	Products        []ProductRequest
	DeliveryAddress DeliveryAddress
	PaymentMethod   string
	PaymentToken    *string
	Promocode       *string
}

type OrderDatabase struct {
	CustomerID    int64
	PaymentMethod string
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
