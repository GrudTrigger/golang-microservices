package order

import (
	"context"

	ordersV1 "github.com/rocker-crm/shared/pkg/openapi/orders/v1"
	"github.com/rocket-crm/order/internal/model"
)

func (r *repository) Update(ctx context.Context, uuid, transactionUuid, paymentMethod, status string) (string, error) {
	order, err := r.GetByUuid(ctx, uuid)
	if err != nil {
		return "", model.ErrOrderNotFound
	}

	order.Status = status
	if transactionUuid != "" {
		order.TransactionUUID = ordersV1.NewOptString(transactionUuid)
	}
	if paymentMethod != "" {
		order.PaymentMethod = ordersV1.NewOptString(paymentMethod)
	}

	_, err = r.db.Exec(ctx, "UPDATE orders SET status=$1, transaction_uuid=$2, payment_method=$3 WHERE id=$4", order.Status, order.TransactionUUID, order.PaymentMethod, uuid)
	if err != nil {
		return "", err
	}

	return order.TransactionUUID.Value, nil
}
