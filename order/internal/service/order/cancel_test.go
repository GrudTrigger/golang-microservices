package order

import (
	"context"

	"github.com/google/uuid"
	"github.com/rocket-crm/order/internal/model"
)

func (s *ServiceSuite) TestCancelOrder() {
	ctx := context.Background()
	orderUuid := uuid.NewString()
	userUuid := uuid.NewString()
	uuidParts := []string{uuid.NewString(), uuid.NewString()}
	transUuid := uuid.NewString()
	var totalPrice float32 = 2000
	orderRepo := model.Order{
		OrderUUID:  orderUuid,
		TotalPrice: totalPrice,
		PartUuids:  uuidParts,
		UserUUID:   userUuid,
		Status:     CANCELLED,
	}

	s.orderRepository.On("GetByUuid", orderUuid).Return(orderRepo, nil)
	s.orderRepository.On("Update", orderRepo.OrderUUID, "", "", CANCELLED).Return(transUuid, nil)
	_, err := s.service.CancelOrder(ctx, orderUuid)
	s.Require().NoError(err)
}
