package v1

import (
	"context"

	paymentV1 "github.com/rocker-crm/shared/pkg/proto/payment/v1"
	"github.com/rocket-crm/payment/internal/converter"
)

func (a *api) PayOrder(ctx context.Context, req *paymentV1.PayOrderRequest) (*paymentV1.PayOrderResponse, error) {
	transactionUuid, err := a.paymentService.PayOrder(ctx, converter.PayOrderToModel(req))
	if err != nil {
		return &paymentV1.PayOrderResponse{}, err
	}
	return &paymentV1.PayOrderResponse{TransactionUuid: transactionUuid}, nil
}
