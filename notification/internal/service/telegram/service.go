package telegram

import (
	"bytes"
	"context"
	"embed"
	"text/template"

	"github.com/rocker-crm/notifacation/internal/client"
	"github.com/rocker-crm/notifacation/internal/model"
	"github.com/rocker-crm/platform/pkg/logger"
	"go.uber.org/zap"
)

const chatID = 123123213 //Узнать чат id бота

//go:embed templates/order_paid_notification.tmpl
var templateOrderPaidFS embed.FS

//go:embed templates/ship_assembled_notification.tmpl
var templateShipAssembled embed.FS

type orderPaidTemplateData struct {
	EventUuid       string
	OrderUuid       string
	UserUuid        string
	PaymentMethod   string
	TransactionUuid string
}

type shipAssembledTemplateData struct {
	EventUuid string
	OrderUuid string
	UserUuid  string
	BuildTime int64
}

type soundValue struct {
	Value bool
}

var orderPaidTemplate = template.Must(template.ParseFS(templateOrderPaidFS, "templates/order_paid_notification.tmpl"))
var shipAssembledTemplate = template.Must(template.ParseFS(templateShipAssembled, "templates/ship_assembled_notification.tmpl"))

type service struct {
	telegramClient client.TelegramClient
}

func NewService(telegramClient client.TelegramClient) *service {
	return &service{
		telegramClient: telegramClient,
	}
}

func (s *service) SendOrderPaidNotification(ctx context.Context, orderEvent model.OrderPaidEvent) error {
	message, err := s.buildOrderPaidMessage(orderEvent)
	if err != nil {
		return err
	}

	err = s.telegramClient.SendMessage(ctx, chatID, message)
	if err != nil {
		return err
	}

	logger.Info(ctx, "Telegram message sent to chat", zap.Int("chat_id", chatID), zap.String("message", message))
	return nil
}

func (s *service) buildOrderPaidMessage(order model.OrderPaidEvent) (string, error) {
	data := orderPaidTemplateData{
		EventUuid: order.EventUuid,
		OrderUuid: order.OrderUuid,
		UserUuid: order.UserUuid,
		PaymentMethod: order.PaymentMethod,
		TransactionUuid: order.TransactionUuid,
	}

	var buf bytes.Buffer
	err := orderPaidTemplate.Execute(&buf, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (s *service) SendShipAssembledNotification(ctx context.Context, shipEvent model.ShipAssembledEvent) error {
	message, err := s.buildShipAssembledMessage(shipEvent)
	if err != nil {
		return err
	}

	err = s.telegramClient.SendMessage(ctx, chatID, message)
	if err != nil {
		return err
	}

	logger.Info(ctx, "Telegram message sent to chat", zap.Int("chat_id", chatID), zap.String("message", message))
	return nil
}

func (s *service) buildShipAssembledMessage(shipAssembled model.ShipAssembledEvent) (string, error) {
	data := shipAssembledTemplateData {
		EventUuid: shipAssembled.EventUuid,
		OrderUuid: shipAssembled.OrderUuid,
		UserUuid: shipAssembled.UserUuid,
		BuildTime: shipAssembled.BuildTime,
	}

	var buf bytes.Buffer
	err := shipAssembledTemplate.Execute(&buf, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

