package order

import (
	"github.com/rocket-crm/order/internal/client/grpc"
	"github.com/rocket-crm/order/internal/repository"
	serviceInterface "github.com/rocket-crm/order/internal/service"
)

type service struct {
	orderRepository repository.OrderRepository
	inventoryClient grpc.InventoryClient
	paymentClient   grpc.PaymentClient
	producer serviceInterface.ProducerService
}

func NewService(orderRepository repository.OrderRepository, inventoryClient grpc.InventoryClient, paymentClient grpc.PaymentClient, producer serviceInterface.ProducerService) *service {
	return &service{
		orderRepository: orderRepository,
		inventoryClient: inventoryClient,
		paymentClient:   paymentClient,
		producer: producer,
	}
}
