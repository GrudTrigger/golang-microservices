package order

import (
	ordersV1 "github.com/rocker-crm/shared/pkg/openapi/orders/v1"
	"github.com/rocket-crm/order/internal/model"
)

func (r *repository) Update(uuid string, transactionUuid string, paymentMethod string, status string) (string, error) {
	order, ok := r.orders[uuid]
	if !ok {
		return "", model.ErrOrderNotFound
	}
	order.Status = status
	if transactionUuid != "" {
		order.TransactionUUID = ordersV1.NewOptString(transactionUuid)

	}

	if paymentMethod != "" {
		order.PaymentMethod = ordersV1.NewOptString(paymentMethod)
	}
	r.orders[uuid] = order
	return order.TransactionUUID.Value, nil
}
