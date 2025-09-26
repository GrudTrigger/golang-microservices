package v1

import (
	"context"
	"errors"
	"net/http"

	ordersV1 "github.com/rocker-crm/shared/pkg/openapi/orders/v1"
	"github.com/rocket-crm/order/internal/model"
)

func (a *api) PayOrder(ctx context.Context, req *ordersV1.PayOrderRequest, params ordersV1.PayOrderParams) (ordersV1.PayOrderRes, error) {
	transactionUuid, err := a.orderService.PayOrder(ctx, req.PaymentMethod, params.OrderUUID)
	if err != nil {
		if errors.Is(err, model.ErrOrderNotFound) {
			return &ordersV1.NotFoundError{
				Code:    http.StatusNotFound,
				Message: "Заказ с uuid " + params.OrderUUID + " не найден!",
			}, nil
		}
		return nil, err
	}
	return &ordersV1.PayOrderResponse{TransactionUUID: transactionUuid}, err
}
