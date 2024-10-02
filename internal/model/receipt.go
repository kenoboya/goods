package model

import "time"

type Receipt struct {
	ReceiptID   int64      `json:"receipt_id" db:"receipt_id"`
	TotalAmount float64    `json:"total_amount" db:"total_amount"`
	ReceiptDate *time.Time `json:"receipt_date" db:"receipt_date"`
}

type ReceiptDetails struct {
	Receipt Receipt
	Order   Order
}
