package converter

import (
	"github.com/rocket-crm/order/internal/model"
	repoModal "github.com/rocket-crm/order/internal/repository/model"
)

func OrderToModal(order repoModal.Order) model.Order {
	return model.Order{
		UserUUID:        order.UserUUID,
		OrderUUID:       order.OrderUUID,
		PaymentMethod:   order.PaymentMethod,
		TransactionUUID: order.TransactionUUID,
		Status:          order.Status,
		TotalPrice:      order.TotalPrice,
		PartUuids:       order.PartUuids,
	}
}
