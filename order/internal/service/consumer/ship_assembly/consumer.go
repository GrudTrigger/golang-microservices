package ship_assembly

import (
	"context"

	"github.com/rocker-crm/platform/pkg/kafka"
	"github.com/rocker-crm/platform/pkg/logger"
	"github.com/rocket-crm/order/internal/repository"
	"go.uber.org/zap"
)

type service struct {
	shipAssemblyConsumer kafka.Consumer
	orderRepository      repository.OrderRepository
}

func NewService(shipAssemblyConsumer kafka.Consumer, orderRepository repository.OrderRepository) *service {
	return &service{
		shipAssemblyConsumer: shipAssemblyConsumer,
		orderRepository:      orderRepository,
	}
}

func (s *service) RunConsumer(ctx context.Context) error {
	logger.Info(ctx, "Starting shipAssemblyConsumer service")
	err := s.shipAssemblyConsumer.Consume(ctx, s.ShipAssembledHandler)
	if err != nil {
		logger.Error(ctx, "Consume from order.assembled topic error", zap.Error(err))
		return err
	}
	return nil
}
