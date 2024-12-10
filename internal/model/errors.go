package model

import "errors"

var (
	ErrNotFoundConfigFile = errors.New("failed to find config file")
	ErrNotFoundEnvFile    = errors.New("failed to load environment file")
	ErrNotFoundCategory   = errors.New("category not found")
	ErrNotFoundProducts   = errors.New("products not found")
	ErrNotFoundProduct    = errors.New("product not found")
	ErrNotFoundCustomer   = errors.New("customer not found")
	ErrNotFoundPromocode  = errors.New("promocode not found")
	ErrNotFoundOrder      = errors.New("order not found")
	ErrEmptyParam         = errors.New("param is empty")
	ErrInvalidRole        = errors.New("invalid role")
	ErrContextIsEmpty     = errors.New("context is empty")
)
