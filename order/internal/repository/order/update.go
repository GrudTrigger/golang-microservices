package order

import (
	ordersV1 "github.com/rocker-crm/shared/pkg/openapi/orders/v1"
	"github.com/rocket-crm/order/internal/model"
)

const PAID = "PAID"

func (r *repository) Update(uuid string, transactionUuid string, paymentMethod string) (string, error) {
	order, ok := r.orders[uuid]
	if !ok {
		return "", model.ErrOrderNotFound
	}
	order.Status = PAID
	order.TransactionUUID = ordersV1.NewOptString(transactionUuid)
	order.PaymentMethod = ordersV1.NewOptString(paymentMethod)

	return order.TransactionUUID.Value, nil
}
