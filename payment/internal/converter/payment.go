package converter

import (
	paymentV1 "github.com/rocker-crm/shared/pkg/proto/payment/v1"
	"github.com/rocket-crm/payment/internal/model"
)

func PayOrderToModel(pay *paymentV1.PayOrderRequest) model.PayOrder {
	return model.PayOrder{
		OrderUuid: pay.OrderUuid,
		UserUuid: pay.UserUuid,
		PaymentMethod: model.PaymentMethod(pay.PaymentMethod),
	}
}