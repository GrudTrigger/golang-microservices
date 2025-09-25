package service

import (
	"context"

	"github.com/rocket-crm/payment/internal/model"
)

type PaymentService interface {
	PayOrder(context.Context, model.PayOrder) (string, error)
}