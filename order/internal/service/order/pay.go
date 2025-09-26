package order

import (
	"context"

	paymentV1 "github.com/rocker-crm/shared/pkg/proto/payment/v1"
	"github.com/rocket-crm/order/internal/model"
)

const PAID = "PAID"

func (s *service) PayOrder(ctx context.Context, paymentMethod string, orderUuid string) (string, error) {
	order, err := s.orderRepository.GetByUuid(orderUuid)
	if err != nil {
		return "", model.ErrOrderNotFound
	}

	resp, err := s.paymentClient.PayOrder(ctx, &paymentV1.PayOrderRequest{UserUuid: order.UserUUID, OrderUuid: order.OrderUUID, PaymentMethod: paymentV1.PaymentMethod(paymentV1.PaymentMethod_value[paymentMethod])})
	if err != nil {
		return "", err
	}
	transactionUuid, err := s.orderRepository.Update(order.OrderUUID, resp.TransactionUuid, paymentMethod)
	if err != nil {
		return "", model.ErrOrderNotFound
	}

	return transactionUuid, nil
}
