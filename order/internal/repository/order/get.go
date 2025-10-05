package order

import (
	"context"
	"database/sql"
	"strconv"

	ordersV1 "github.com/rocker-crm/shared/pkg/openapi/orders/v1"
	"github.com/rocket-crm/order/internal/model"
	"github.com/rocket-crm/order/internal/repository/converter"
	repoModal "github.com/rocket-crm/order/internal/repository/model"
)

func (r *repository) GetByUuid(ctx context.Context, uuid string) (model.Order, error) {
	var order repoModal.Order
	var transactionUUID sql.NullString
	var paymentMethod sql.NullString
	idInt, err := strconv.Atoi(uuid)
	if err != nil {
		return model.Order{}, err
	}
	row := r.db.QueryRow(ctx, "SELECT id, part_uuid, total_price, transaction_uuid, payment_method, status, user_uuid FROM orders WHERE id=$1", idInt)
	err = row.Scan(&order.OrderUUID, &order.PartUuids, &order.TotalPrice, &transactionUUID, &paymentMethod, &order.Status, &order.UserUUID)
	if err != nil {
		return model.Order{}, err
	}
	order.TransactionUUID = ordersV1.OptString{
		Value: transactionUUID.String,
		Set:   transactionUUID.Valid,
	}
	order.PaymentMethod = ordersV1.OptString{
		Value: paymentMethod.String,
		Set:   paymentMethod.Valid,
	}
	return converter.OrderToModal(order), nil
}
