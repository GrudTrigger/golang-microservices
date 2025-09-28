package service

import (
	"context"

	ordersV1 "github.com/rocker-crm/shared/pkg/openapi/orders/v1"
	"github.com/rocket-crm/order/internal/model"
)

type OrderService interface {
	CreateOrder(ctx context.Context, data model.CreateOrder) (model.ResponseCreateOrder, error)
	GetOrderByUuid(ctx context.Context, orderUuid string) (model.Order, error)
	PayOrder(ctx context.Context, paymentMethod, orderUuid string) (string, error)
	CancelOrder(ctx context.Context, orderUuid string) (ordersV1.CancelOrderRes, error)
}
