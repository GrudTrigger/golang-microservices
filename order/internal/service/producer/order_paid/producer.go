package order_paid

import (
	"context"

	"github.com/rocker-crm/platform/pkg/kafka"
	"github.com/rocker-crm/platform/pkg/logger"
	eventsV1 "github.com/rocker-crm/shared/pkg/proto/events/v1"
	"github.com/rocket-crm/order/internal/model"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

type service struct {
	orderPaidProducer kafka.Producer
}

func NewService(orderPaidProducer kafka.Producer) *service {
	return &service{orderPaidProducer: orderPaidProducer}
}

func (s *service) ProducerOrderPaidRecorder(ctx context.Context, event model.OrderPaidEvent) error {
	msg := &eventsV1.OrderPaidRecorder{
		EventUuid:       event.EventUuid,
		OrderUuid:       event.OrderUuid,
		UserUuid:        event.UserUuid,
		PaymentMethod:   event.PaymentMethod,
		TransactionUuid: event.TransactionUuid,
	}

	payload, err := proto.Marshal(msg)
	if err != nil {
		logger.Error(ctx, "failed to marshal ProducerOrderPaidRecorder", zap.Error(err))
		return err
	}

	err = s.orderPaidProducer.Send(ctx, []byte(event.OrderUuid), payload)
	if err != nil {
		logger.Error(ctx, "failed to publish ProducerOrderPaidRecorder", zap.Error(err))
		return err
	}

	return nil
}
