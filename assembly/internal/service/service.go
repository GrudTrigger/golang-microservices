package service

import (
	"context"

	"github.com/rocker-crm/assembly/internal/model"
)

type ConsumerService interface {
	RunConsumer(ctx context.Context) error
}

type ProducerService interface {
	ProduceShipAssembledRecorded(ctx context.Context, event model.ShipAssembledEvent) error
}
