package order

import (
	"github.com/google/uuid"
	"github.com/rocket-crm/order/internal/model"
	"github.com/rocket-crm/order/internal/repository/converter"
	repoModal "github.com/rocket-crm/order/internal/repository/model"
)

const PendingPayment = "PENDING_PAYMENT"

func (r *repository) Create(req model.CreateOrder, totalPrice float32) model.Order {
	order := repoModal.Order{
		OrderUUID:  uuid.NewString(),
		UserUUID:   req.UserUUID,
		PartUuids:  req.PartUuids,
		TotalPrice: totalPrice,
		Status:     PendingPayment,
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	r.orders[order.OrderUUID] = order
	return converter.OrderToModal(order)
}
