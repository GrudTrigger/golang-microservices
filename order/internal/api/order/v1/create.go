package v1

import (
	"context"

	ordersV1 "github.com/rocker-crm/shared/pkg/openapi/orders/v1"
	"github.com/rocket-crm/order/internal/converter"
)

func (a *api) CreateOrder(ctx context.Context, req *ordersV1.CreateOrderRequest) (ordersV1.CreateOrderRes, error) {
	order, err := a.orderService.CreateOrder(ctx, converter.CreateOrderToModel(req))
	if err != nil {
		return nil, err
	}
	return &ordersV1.CreateOrderResponse{OrderUUID: order.UUID, TotalPrice: order.TotalPrice}, nil
}
