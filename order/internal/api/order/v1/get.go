package v1

import (
	"context"
	"net/http"

	"github.com/go-faster/errors"
	ordersV1 "github.com/rocker-crm/shared/pkg/openapi/orders/v1"
	"github.com/rocket-crm/order/internal/converter"
	"github.com/rocket-crm/order/internal/model"
)

func (a *api) GetOrderByUuid(ctx context.Context, params ordersV1.GetOrderByUuidParams) (ordersV1.GetOrderByUuidRes, error) {
	order, err := a.orderService.GetOrderByUuid(ctx, params.OrderUUID)
	if err != nil {
		if errors.Is(err, model.ErrOrderNotFound) {
			return &ordersV1.NotFoundError{
				Code:    http.StatusNotFound,
				Message: "Заказ с uuid " + params.OrderUUID + " не найден!",
			}, nil
		}
		return nil, err
	}
	return converter.OrderModelToOrder(order), nil
}
