package order

import (
	"github.com/google/uuid"
	"github.com/rocket-crm/order/internal/model"
)

func (s *ServiceSuite) TestGetOrderByUuidSuccess() {
	orderUuid := uuid.NewString()
	var totalPrice float32 = 3700
	uuidParts := []string{uuid.NewString(), uuid.NewString()}
	order := model.Order{
		OrderUUID:  orderUuid,
		TotalPrice: totalPrice,
		PartUuids:  uuidParts,
		UserUUID:   uuid.NewString(),
		Status:     PAID,
	}
	s.orderRepository.On("GetByUuid", orderUuid).Return(order, nil)
	o, err := s.service.GetOrderByUuid(s.ctx, orderUuid)
	s.Require().NoError(err)
	s.Equal(o, order)
}

func (s *ServiceSuite) TestGetOrderByUuidError() {
	orderUuid := uuid.NewString()
	s.orderRepository.On("GetByUuid", orderUuid).Return(model.Order{}, model.ErrOrderNotFound)
	resp, err := s.service.GetOrderByUuid(s.ctx, orderUuid)
	s.Require().ErrorIs(err, model.ErrOrderNotFound)
	s.Require().Empty(resp)
}
