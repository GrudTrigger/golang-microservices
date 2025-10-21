package order_consumer

import (
	"context"

	"github.com/rocker-crm/platform/pkg/kafka"
	"github.com/rocker-crm/platform/pkg/logger"
	"go.uber.org/zap"
)

type service struct {
	orderRecorderConsumer kafka.Consumer
}

func NewConsumerService(orderRecorderConsumer kafka.Consumer) *service {
	return &service{
		orderRecorderConsumer: orderRecorderConsumer,
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
