package order

import (
	"context"

	"github.com/google/uuid"
	"github.com/rocket-crm/order/internal/model"
	"github.com/rocket-crm/order/internal/repository/converter"
	repoModal "github.com/rocket-crm/order/internal/repository/model"
)

const PendingPayment = "PENDING_PAYMENT"

func (r *repository) Create(ctx context.Context, req model.CreateOrder, totalPrice float32) (model.Order, error) {
	order := repoModal.Order{
		OrderUUID:  uuid.NewString(),
		UserUUID:   req.UserUUID,
		PartUuids:  req.PartUuids,
		TotalPrice: totalPrice,
		Status:     PendingPayment,
	}

	_, err := r.db.Exec(ctx, "INSERT INTO orders(id, part_uuid, total_price, status, user_uuid) VALUES($1, $2, $3, $4, $5)", order.OrderUUID, order.PartUuids, order.TotalPrice, order.Status, order.UserUUID)
	if err != nil {
		return model.Order{}, err
	}

	return converter.OrderToModal(order), nil
}
