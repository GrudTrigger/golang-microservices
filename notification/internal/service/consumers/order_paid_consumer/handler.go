package order_paid_consumer

import (
	"context"

	"github.com/rocker-crm/notifacation/internal/converter/decoder"
	"github.com/rocker-crm/platform/pkg/kafka"
	"github.com/rocker-crm/platform/pkg/logger"
	"go.uber.org/zap"
)

func (s *service) OrderHandler(ctx context.Context, msg kafka.Message) error {
	event, err := decoder.Decode(msg.Value)
	if err != nil {
		logger.Error(ctx, "Failed to decode OrderPaidRecorded", zap.Error(err))
		return err
	}
	err = s.telegramService.SendOrderPaidNotification(ctx, event)
	if err != nil {
		logger.Error(ctx, "Failed to send message to telegram", zap.Error(err))
		return err
	}
	return nil
}
