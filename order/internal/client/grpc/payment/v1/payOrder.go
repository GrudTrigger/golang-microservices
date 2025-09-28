package v1

import (
	"context"

	paymentV1 "github.com/rocker-crm/shared/pkg/proto/payment/v1"
	"github.com/rocket-crm/order/internal/model"
)

func (c *client) PayOrder(ctx context.Context, req model.RequestPay) (string, error) {
	res, err := c.generatedClient.PayOrder(ctx, &paymentV1.PayOrderRequest{
		UserUuid:      req.UserUuid,
		PaymentMethod: paymentV1.PaymentMethod(req.PaymentMethod),
		OrderUuid:     req.OrderUuid,
	})
	if err != nil {
		return "", err
	}
	return res.TransactionUuid, nil
}
