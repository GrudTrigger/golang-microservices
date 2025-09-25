package v1

import (
	paymentV1 "github.com/rocker-crm/shared/pkg/proto/payment/v1"
	"github.com/rocket-crm/payment/internal/service"
)

type api struct {
	paymentV1.UnimplementedPaymentServiceServer
	paymentService service.PaymentService
}

func NewAPI(paymentService service.PaymentService) *api {
	return &api{paymentService: paymentService}
}
