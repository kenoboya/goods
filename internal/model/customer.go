package model

type Customer struct {
	CustomerID       int64   `db:"customer_id"`
	UserID           *string `json:"user_id" db:"user_id"`
	CustomerFullName string  `json:"customer_full_name" db:"customer_full_name"`
	CustomerPhone    string  `json:"customer_phone" db:"customer_phone"`
}

type RequestCustomer struct {
	CustomerFullName string
	CustomerEmail    string
	CustomerPhone    string
}
