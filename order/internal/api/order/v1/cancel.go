package v1

import (
	"context"
	"errors"
	"net/http"

	ordersV1 "github.com/rocker-crm/shared/pkg/openapi/orders/v1"
	"github.com/rocket-crm/order/internal/model"
)

func (a *api) CancelOrder(ctx context.Context, params ordersV1.CancelOrderParams) (ordersV1.CancelOrderRes, error) {
	res, err := a.orderService.CancelOrder(ctx, params.OrderUUID)
	if err != nil {
		if errors.Is(err, model.ErrOrderNotFound) {
			return &ordersV1.NotFoundError{
				Code:    http.StatusNotFound,
				Message: "Заказ с uuid " + params.OrderUUID + " не найден!",
			}, nil
		}
		return nil, err
	}
	return res, nil
}
