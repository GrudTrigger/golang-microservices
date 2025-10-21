package model

type OrderPaidEvent struct {
	EventUuid       string
	OrderUuid       string
	UserUuid        string
	PaymentMethod   string
	TransactionUuid string
}
