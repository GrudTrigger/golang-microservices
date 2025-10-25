package telegram

import (
	"embed"
	"text/template"

	"github.com/rocker-crm/notifacation/internal/client"
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
