package order

import (
	"context"

	"github.com/rocket-crm/order/internal/model"
	"github.com/rocket-crm/order/internal/repository/converter"
	repoModal "github.com/rocket-crm/order/internal/repository/model"
)

func (r *repository) GetByUuid(ctx context.Context, uuid string) (model.Order, error) {
	var order repoModal.Order
	row := r.db.QueryRow(ctx, "SELECT id, part_uuid, total_price, transaction_uuid, payment_method, status, user_uuid FROM orders WHERE id=$1", uuid)
	err := row.Scan(&order)
	if err != nil {
		return model.Order{}, err
	}
	return converter.OrderToModal(order), nil
}
