package order

import (
	inventoryV1 "github.com/rocker-crm/shared/pkg/proto/inventory/v1"
	paymentV1 "github.com/rocker-crm/shared/pkg/proto/payment/v1"
	"github.com/rocket-crm/order/internal/repository"
)

type service struct {
	orderRepository repository.OrderRepository
	inventoryClient inventoryV1.InventoryServiceClient
	paymentClient   paymentV1.PaymentServiceClient
}

func NewService(orderRepository repository.OrderRepository, inventoryClient inventoryV1.InventoryServiceClient, paymentClient paymentV1.PaymentServiceClient) *service {
	return &service{
		orderRepository: orderRepository,
		inventoryClient: inventoryClient,
		paymentClient:   paymentClient,
	}
}
