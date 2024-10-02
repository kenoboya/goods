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
	CustomerID       *string `json:"customer_id" db:"customer_id"`
	CustomerFullName string  `json:"customer_full_name" db:"customer_full_name"`
	CustomerPhone    string  `json:"customer_phone" db:"customer_phone"`
}

type Order struct {
	OrderID     int64      `json:"order_id" db:"order_id"`
	OrderDate   *time.Time `json:"order_date" db:"order_date"`
	ShipperID   *string    `json:"shipper_id" db:"shipper_id"`
	AcceptedAt  *time.Time `json:"accepted_at" db:"accepted_at"`
	DeliveredAt *time.Time `json:"delivered_at" db:"delivered_at"`
	Customer
	DeliveryAddress
}

type OrderDetails struct {
	Order     Order     `json:"order"`
	Product   Product   `json:"product"`
	Promocode Promocode `json: "promocode"`
	Quantity  int8      `json:"quantity" db:"quantity"`
}
