package model

import ordersV1 "github.com/rocker-crm/shared/pkg/openapi/orders/v1"

type Order struct {
	OrderUUID       string             `json:"id"`
	UserUUID        string             `json:"user_uuid"`
	PartUuids       []string           `json:"part_uuid"`
	TotalPrice      float32            `json:"total_price"`
	TransactionUUID ordersV1.OptString `json:"transaction_uuid"`
	PaymentMethod   ordersV1.OptString `json:"payment_method"`
	Status          string             `json:"status"`
}

type CreateOrder struct {
	UserUUID  string
	PartUuids []string
}
