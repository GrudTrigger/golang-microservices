package model

import ordersV1 "github.com/rocker-crm/shared/pkg/openapi/orders/v1"

type Order struct {
	OrderUUID       string
	UserUUID        string
	PartUuids       []string
	TotalPrice      float32
	TransactionUUID ordersV1.OptString
	PaymentMethod   ordersV1.OptString
	Status          string
}

type CreateOrder struct {
	UserUUID  string
	PartUuids []string
}
