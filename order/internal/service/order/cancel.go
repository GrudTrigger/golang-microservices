package order

import (
	"context"

	ordersV1 "github.com/rocker-crm/shared/pkg/openapi/orders/v1"
)

const CANCELLED = "CANCELLED"

func (s *service) CancelOrder(ctx context.Context, orderUuid string) (ordersV1.CancelOrderRes, error) {
	order, err := s.orderRepository.GetByUuid(orderUuid)
	if err != nil {
		return nil, err
	}

	if order.Status == PAID {
		return &ordersV1.ConflictError{}, nil
	}

	order.Status = CANCELLED
	return &ordersV1.CancelOrderNoContent{}, nil
}
