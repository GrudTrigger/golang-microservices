package order

import (
	"github.com/rocket-crm/order/internal/model"
	"github.com/rocket-crm/order/internal/repository/converter"
)

func (r *repository) GetByUuid(uuid string) (model.Order, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	order, ok := r.orders[uuid]
	if !ok {
		return model.Order{}, model.ErrOrderNotFound
	}
	return converter.OrderToModal(order), nil
}
