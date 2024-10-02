package model

import "time"

type Promocode struct {
	PromocodeID string     `json:"promocode_id" db:"promocode_id"`
	Discount    *float64   `json:"discount" db:"discount"`
	StartAt     *time.Time `json:"start_at" db:"start_at"`
	FinishAt    *time.Time `json:"finish_at" db:"finish_at"`
}
