package grpc

import (
	"context"

	"github.com/rocket-crm/order/internal/model"
)

type InventoryClient interface {
	ListParts(ctx context.Context, filter model.PartsFilter) ([]model.Part, error)
}

type PaymentClient interface {
	PayOrder(ctx context.Context, req model.RequestPay) (string, error)
}
