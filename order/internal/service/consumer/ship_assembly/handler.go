package ship_assembly

import (
	"context"

	"github.com/rocker-crm/platform/pkg/kafka"
	"github.com/rocker-crm/platform/pkg/logger"
	"github.com/rocket-crm/order/internal/converter/decoder"
	"go.uber.org/zap"
)

func (s *service) ShipAssembledHandler(ctx context.Context, msg kafka.Message) error {
	event, err := decoder.Decode(msg.Value)
	if err != nil {
		logger.Error(ctx, "Failed to decode ShipAssembledRecorded", zap.Error(err))
		return err
	}

	_, err = s.orderRepository.Update(ctx, event.OrderUuid, "", "", "COMPLETED")
	if err != nil {
		logger.Error(ctx, "Failed to update status order in handler consumer", zap.Error(err))
		return err
	}

	return nil
}
