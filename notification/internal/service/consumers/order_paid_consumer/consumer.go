package order_paid_consumer

import (
	"context"

	notificationService "github.com/rocker-crm/notifacation/internal/service"
	"github.com/rocker-crm/platform/pkg/kafka"
	"github.com/rocker-crm/platform/pkg/logger"
	"go.uber.org/zap"
)

type service struct {
	orderRecorderConsumer kafka.Consumer
	telegramService       notificationService.TelegramService
}

func NewOrderPaidConsumerService(orderRecorderConsumer kafka.Consumer, telegramService notificationService.TelegramService) *service {
	return &service{
		orderRecorderConsumer: orderRecorderConsumer,
		telegramService:       telegramService,
	}
}

func (s *service) RunConsumer(ctx context.Context) error {
	logger.Info(ctx, "Starting orderRecorderConsumer service")

	err := s.orderRecorderConsumer.Consume(ctx, s.OrderHandler)
	if err != nil {
		logger.Error(ctx, "Consume from order.paid service:notification topic error", zap.Error(err))
		return err
	}
	return nil
}
