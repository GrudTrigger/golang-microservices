package order

import (
	"context"

	"github.com/rocket-crm/order/internal/model"
)

func (s *service) GetOrderByUuid(ctx context.Context, orderUuid string) (model.Order, error) {
	order, err := s.orderRepository.GetByUuid(orderUuid)
	if err != nil {
		return model.Order{}, model.ErrOrderNotFound
	}
	return order, nil
}
