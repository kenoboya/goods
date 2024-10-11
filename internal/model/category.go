package model

type Category struct {
	CategoryID int8 `json:"category_id" db:"category_id"`
	CreateCategoryRequest
}

type CreateCategoryRequest struct {
	CategoryName string  `json:"category_name" db:"category_name"`
	Description  string  `json:"description" db:"description"`
	Image        *string `json:"image" db:"image"`
}
