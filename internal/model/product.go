package model

type Category struct {
	CategoryID   int8    `json:"category_id" db:"category_id"`
	CategoryName string  `json:"category_name" db:"category_name"`
	Description  string  `json:"description" db:"description"`
	Image        *string `json:"image" db:"image"`
}

type Product struct {
	ProductID       int     `json:"product_id" db:"product_id"`
	ProductName     int     `json:"product_name" db:"product_name"`
	SupplierID      int8    `json:"supplier_id" db:"supplier_id"`
	CategoryID      int8    `json:"category_id" db:"category_id"`
	UnitPrice       float64 `json:"unit_price" db:"unit_price"`
	Stock           bool    `json:"in_stock" db:"in_stock"`
	Discount        float64 `json:"discount" db:"discount"`
	QuantityPerUnit *string `json:"quantity_per_unit" db:"quantity_per_unit"`
	Weight          *string `json:"weight" db:"weight"`
	Image           *string `json:"image" db:"image"`
}

type ProductRequest struct {
	ProductID int
	Quantity  int8
}

type ProductResponse struct {
	Product  Product
	Quantity int8
}
