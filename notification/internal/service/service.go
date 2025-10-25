package service

import (
	"context"

	"github.com/rocker-crm/notifacation/internal/model"
)

type ConsumerService interface {
	RunConsumer(ctx context.Context) error
}

type TelegramService interface {
	SendOrderPaidNotification(ctx context.Context, orderEvent model.OrderPaidEvent) error
	SendShipAssembledNotification(ctx context.Context, shipEvent model.ShipAssembledEvent) error
}
