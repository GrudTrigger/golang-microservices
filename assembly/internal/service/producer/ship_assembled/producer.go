package ship_assembled

import (
	"context"

	"github.com/rocker-crm/assembly/internal/model"
	"github.com/rocker-crm/platform/pkg/kafka"
	"github.com/rocker-crm/platform/pkg/logger"
	eventsV1 "github.com/rocker-crm/shared/pkg/proto/events/v1"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

type service struct {
	shipAssembledProducer kafka.Producer
}

func NewService(shipAssembledProducer kafka.Producer) *service {
	return &service{shipAssembledProducer: shipAssembledProducer}
}

func (s *service) ProduceShipAssembledRecorded(ctx context.Context, event model.ShipAssembledEvent) error {
	msg := &eventsV1.ShipAssembledRecorder{
		EventUuid:    event.EventUuid,
		OrderUuid:    event.OrderUuid,
		UserUuid:     event.UserUuid,
		BuildTimeSec: event.BuildTime,
	}

	payload, err := proto.Marshal(msg)
	if err != nil {
		logger.Error(ctx, "failed to marshal ProduceShipAssembledRecorded", zap.Error(err))
		return err
	}

	err = s.shipAssembledProducer.Send(ctx, []byte(event.OrderUuid), payload)
	if err != nil {
		logger.Error(ctx, "failed to publish ProduceShipAssembledRecorded", zap.Error(err))
		return err
	}

	return nil
}
