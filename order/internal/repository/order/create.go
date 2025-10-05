package order

import (
	"context"

	"github.com/rocket-crm/order/internal/model"
	"github.com/rocket-crm/order/internal/repository/converter"
	repoModal "github.com/rocket-crm/order/internal/repository/model"
)

const PendingPayment = "PENDING_PAYMENT"

func (r *repository) Create(ctx context.Context, req model.CreateOrder, totalPrice float32) (model.Order, error) {
	order := repoModal.Order{
		UserUUID:   req.UserUUID,
		PartUuids:  req.PartUuids,
		TotalPrice: totalPrice,
		Status:     PendingPayment,
	}

	var orderId string
	row := r.db.QueryRow(ctx, "INSERT INTO orders(part_uuid, total_price, status, user_uuid) VALUES($1, $2, $3, $4) RETURNING id", order.PartUuids, order.TotalPrice, order.Status, order.UserUUID)
	err := row.Scan(&orderId)
	if err != nil {
		return model.Order{}, err
	}
	order.OrderUUID = orderId
	return converter.OrderToModal(order), nil
}
