package payment

import (
	"testing"

	"github.com/rocket-crm/payment/internal/repository/mocks"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite

	paymentRepository *mocks.PaymentRepository

	service *service
}

func (s *ServiceSuite) SetupTest() {
	s.paymentRepository = mocks.NewPaymentRepository(s.T())

	s.service = NewService(s.paymentRepository)
}

func (s *ServiceSuite) TearDownTest() {}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}
