package model

type Supplier struct {
	SupplierID   int8    `json:"supplier_id" db:"supplier_id"`
	CompanyName  string  `json:"company_name" db:"company_name"`
	ContactName  *string `json:"contact_name" db:"contact_name"`
	ContactTitle *string `json:"contact_title" db:"contact_title"`
	Address      *string `json:"address" db:"address"`
	Phone        *string `json:"phone" db:"phone"`
	Contract     *string `json:"contract" db:"contract"`
}
