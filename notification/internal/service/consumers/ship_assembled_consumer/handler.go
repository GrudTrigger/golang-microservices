package ship_assembled_consumer

import (
	"context"
	"fmt"

	"github.com/rocker-crm/notifacation/internal/converter/decoder"
	"github.com/rocker-crm/platform/pkg/kafka"
	"github.com/rocker-crm/platform/pkg/logger"
	"go.uber.org/zap"
)

func (s *service) ShipAssembledHandler(ctx context.Context, msg kafka.Message) error {
	event, err := decoder.DecodeShipAssembled(msg.Value)
	if err != nil {
		logger.Error(ctx, "Failed to decode OrderPaidRecorded", zap.Error(err))
		return err
	}
	fmt.Println(event)
	return nil
}
