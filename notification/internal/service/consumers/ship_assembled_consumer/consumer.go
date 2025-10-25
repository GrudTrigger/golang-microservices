package ship_assembled_consumer

import (
	"context"

	"github.com/rocker-crm/platform/pkg/kafka"
	"github.com/rocker-crm/platform/pkg/logger"
	"go.uber.org/zap"
)

type service struct {
	shipRecorderConsumer kafka.Consumer
}

func NewShipAssembledConsumerService(shipRecorderConsumer kafka.Consumer) *service {
	return &service{
		shipRecorderConsumer: shipRecorderConsumer,
	}
}

func (s *service) RunConsumer(ctx context.Context) error {
	logger.Info(ctx, "Starting shipAssembledConsumer service")
	err := s.shipRecorderConsumer.Consume(ctx, s.ShipAssembledHandler)
	if err != nil {
		logger.Error(ctx, "Consume from ship.assembled service:notification topic error", zap.Error(err))
		return err
	}
	return nil
}
