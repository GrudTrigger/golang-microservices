package payment

import (
	"errors"

	"github.com/google/uuid"
	"github.com/rocket-crm/payment/internal/model"
)

func (s *ServiceSuite) TestPayOrderSuccess() {
	payOrder := model.PayOrder{OrderUuid: uuid.NewString(), UserUuid: uuid.NewString(), PaymentMethod: 1}

	s.paymentRepository.On("PayOrder", s.ctx, payOrder).Return(uuid.NewString(), nil)
	transactionUuid, err := s.service.PayOrder(s.ctx, payOrder)
	s.Require().NoError(err)
	s.Require().NotEmpty(transactionUuid)
}

func (s *ServiceSuite) TestPayOrderWithError() {
	payOrder := model.PayOrder{OrderUuid: uuid.NewString(), UserUuid: uuid.NewString(), PaymentMethod: 1}
	s.paymentRepository.On("PayOrder", s.ctx, payOrder).Return("", errors.New("fail on generate uuid"))
	transactionUuid, err := s.service.PayOrder(s.ctx, payOrder)
	s.Require().Error(err)
	s.Require().Empty(transactionUuid)
}
