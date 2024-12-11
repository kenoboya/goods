package model

const (
	RoleAdmin    = "admin"
	RoleCustomer = "customer"
	RoleShipper  = "shipper"
)

type UserData struct {
	Fullname string
	Email    string
	Phone    string
}
