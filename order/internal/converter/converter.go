package converter

import (
	ordersV1 "github.com/rocker-crm/shared/pkg/openapi/orders/v1"
	"github.com/rocket-crm/order/internal/model"
)

func CreateOrderToModel(c *ordersV1.CreateOrderRequest) model.CreateOrder {
	return model.CreateOrder{UserUUID: c.UserUUID, PartUuids: c.PartUuids}
}

func OrderModelToOrder(order model.Order) *ordersV1.Order {
	return &ordersV1.Order{
		OrderUUID:       order.OrderUUID,
		TotalPrice:      order.TotalPrice,
		PartUuids:       order.PartUuids,
		UserUUID:        order.UserUUID,
		TransactionUUID: order.TransactionUUID,
		Status:          order.Status,
		PaymentMethod:   order.PaymentMethod,
	}
}

func Ptr[T any](v T) *T {
	return &v
}
