package order_consumer

import (
	"context"

	serviceInterface "github.com/rocker-crm/assembly/internal/service"
	"github.com/rocker-crm/platform/pkg/kafka"
	"github.com/rocker-crm/platform/pkg/logger"
	"go.uber.org/zap"
)

type service struct {
	orderRecorderConsumer kafka.Consumer
	shipProducer          serviceInterface.ProducerService
}

func NewConsumerService(orderRecorderConsumer kafka.Consumer, shipProducer serviceInterface.ProducerService) *service {
	return &service{
		orderRecorderConsumer: orderRecorderConsumer,
		shipProducer:          shipProducer,
	}
}

func (s *service) RunConsumer(ctx context.Context) error {
	logger.Info(ctx, "Starting orderRecorderConsumer service")

	err := s.orderRecorderConsumer.Consume(ctx, s.OrderHandler)
	if err != nil {
		logger.Error(ctx, "Consume from ufo.recorded topic error", zap.Error(err))
		return err
	}

	return nil
}
