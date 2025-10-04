package order

import (
	"context"

	paymentV1 "github.com/rocker-crm/shared/pkg/proto/payment/v1"
	"github.com/rocket-crm/order/internal/model"
)

func (s *service) PayOrder(ctx context.Context, paymentMethod, orderUuid string) (string, error) {
	order, err := s.orderRepository.GetByUuid(ctx, orderUuid)
	if err != nil {
		return "", model.ErrOrderNotFound
	}

	uuid, err := s.paymentClient.PayOrder(ctx, model.RequestPay{UserUuid: order.UserUUID, OrderUuid: order.OrderUUID, PaymentMethod: model.PaymentMethod(paymentV1.PaymentMethod_value[paymentMethod])})
	if err != nil {
		return "", err
	}
	transactionUuid, err := s.orderRepository.Update(ctx, order.OrderUUID, uuid, paymentMethod, "PAID")
	if err != nil {
		return "", model.ErrOrderNotFound
	}

	return transactionUuid, nil
}
