package model

type PaymentMethod int32

type PayOrder struct {
	OrderUuid     string
	UserUuid      string
	PaymentMethod PaymentMethod
}
