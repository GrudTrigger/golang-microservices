package repository

import (
	"github.com/rocket-crm/order/internal/model"
)

type OrderRepository interface {
	Create(req model.CreateOrder, totalPrice float32) model.Order
	GetByUuid(uuid string) (model.Order, error)
	Update(uuid string, transactionUuid string, paymentMethod string, status string) (string, error)
}
