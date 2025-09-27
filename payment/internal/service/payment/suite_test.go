package payment

import (
	"testing"

	"github.com/rocket-crm/payment/internal/repository/mocks"
	"github.com/stretchr/testify/suite"
	"golang.org/x/net/context"
)

type ServiceSuite struct {
	suite.Suite

	ctx context.Context

	paymentRepository *mocks.PaymentRepository

	service *service
}

func (s *ServiceSuite) SetupTest() {
	s.ctx = context.Background()

	s.paymentRepository = mocks.NewPaymentRepository(s.T())

	s.service = NewService(s.paymentRepository)
}

func (s *ServiceSuite) TearDownTest() {}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}
