package order_consumer

import (
	"context"
	"fmt"
	"time"

	"github.com/rocker-crm/assembly/internal/converter/kafka/decoder"
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
	time.Sleep(time.Second * 10)
	fmt.Println(event)
	return nil
}
