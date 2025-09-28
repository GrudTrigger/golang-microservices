package repository

import (
	"context"

	"github.com/rocket-crm/payment/internal/model"
)

type PaymentRepository interface {
	PayOrder(context.Context, model.PayOrder) (string, error)
}
