package model

type PaymentMethod int32

type RequestPay struct {
	OrderUuid string
	UserUuid  string
	PaymentMethod
}
