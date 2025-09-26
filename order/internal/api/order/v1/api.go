package v1

import (
	"context"
	"net/http"

	ordersV1 "github.com/rocker-crm/shared/pkg/openapi/orders/v1"
	"github.com/rocket-crm/order/internal/service"
)

type api struct {
	orderService service.OrderService
}

func NewAPI(orderService service.OrderService) *api {
	return &api{orderService: orderService}
}

func (a *api) NewError(_ context.Context, err error) *ordersV1.GenericErrorStatusCode {
	return &ordersV1.GenericErrorStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: ordersV1.GenericError{
			Code:    ordersV1.NewOptInt(http.StatusInternalServerError),
			Message: ordersV1.NewOptString(err.Error()),
		},
	}
}
