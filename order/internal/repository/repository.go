package repository

import (
	"context"

	"github.com/rocket-crm/order/internal/model"
)

type OrderRepository interface {
	Create(ctx context.Context, req model.CreateOrder, totalPrice float32) (model.Order, error)
	GetByUuid(ctx context.Context, uuid string) (model.Order, error)
	Update(ctx context.Context, uuid, transactionUuid, paymentMethod, status string) (string, error)
}
