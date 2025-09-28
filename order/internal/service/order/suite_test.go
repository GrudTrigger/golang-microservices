package order

import (
	"context"
	"testing"

	mocksClient "github.com/rocket-crm/order/internal/client/grpc/mocks"
	mocksRepo "github.com/rocket-crm/order/internal/repository/mocks"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite

	ctx context.Context

	orderRepository *mocksRepo.OrderRepository
	inventoryClient *mocksClient.InventoryClient
	paymentClient   *mocksClient.PaymentClient

	service *service
}

func (s *ServiceSuite) SetupTest() {
	s.ctx = context.Background()

	s.orderRepository = mocksRepo.NewOrderRepository(s.T())
	s.inventoryClient = mocksClient.NewInventoryClient(s.T())
	s.paymentClient = mocksClient.NewPaymentClient(s.T())

	s.service = NewService(s.orderRepository, s.inventoryClient, s.paymentClient)
}

func (s *ServiceSuite) TearDownTest() {}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}
