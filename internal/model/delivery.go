package model

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