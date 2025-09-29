package order

import (
	"context"

	"github.com/google/uuid"
	"github.com/rocket-crm/order/internal/model"
)

func (s *ServiceSuite) TestPaySuccess() {
	ctx := context.Background()
	orderUuid := uuid.NewString()
	uuidParts := []string{uuid.NewString(), uuid.NewString()}
	tranUuid := uuid.NewString()
	userUuid := uuid.NewString()
	var totalPrice float32 = 2000
	orderRepo := model.Order{
		OrderUUID:  orderUuid,
		TotalPrice: totalPrice,
		PartUuids:  uuidParts,
		UserUUID:   userUuid,
		Status:     PAID,
	}
	s.orderRepository.On("GetByUuid", orderUuid).Return(orderRepo, nil)

	reqPay := model.RequestPay{
		OrderUuid:     orderUuid,
		UserUuid:      userUuid,
		PaymentMethod: 0,
	}
	s.paymentClient.On("PayOrder", ctx, reqPay).Return(tranUuid, nil)

	s.orderRepository.On("Update", orderUuid, tranUuid, "CARD", "PAID").Return(tranUuid, nil)

	res, err := s.service.PayOrder(ctx, "CARD", orderUuid)
	s.Require().NoError(err)
	s.Require().Equal(res, tranUuid)
}

func (s *ServiceSuite) TestPayError() {
	ctx := context.Background()
	orderUuid := uuid.NewString()

	s.orderRepository.On("GetByUuid", orderUuid).Return(model.Order{}, model.ErrOrderNotFound)
	res, err := s.service.PayOrder(ctx, "CARD", orderUuid)
	s.Require().ErrorIs(err, model.ErrOrderNotFound)
	s.Equal(res, "")
}
