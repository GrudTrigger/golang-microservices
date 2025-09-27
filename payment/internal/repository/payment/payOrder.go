package payment

import (
	"context"

	"github.com/google/uuid"
	"github.com/rocket-crm/payment/internal/model"
)

func (r *repository) PayOrder(ctx context.Context, payOrder model.PayOrder) (string, error) {
	tranUuid := uuid.NewString()
	return tranUuid, nil
}
