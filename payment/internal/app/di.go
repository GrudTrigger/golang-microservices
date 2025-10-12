package app

import (
	paymentV1 "github.com/rocker-crm/shared/pkg/proto/payment/v1"
	paymentAPIV1 "github.com/rocket-crm/payment/internal/api/payment/v1"
	"github.com/rocket-crm/payment/internal/repository"
	"github.com/rocket-crm/payment/internal/repository/payment"
	"github.com/rocket-crm/payment/internal/service"
	paymentService "github.com/rocket-crm/payment/internal/service/payment"
)

type diContainer struct {
	paymentAPIV1      paymentV1.PaymentServiceServer
	paymentRepository repository.PaymentRepository
	paymentService    service.PaymentService
}

func NewDiContainer() *diContainer {
	return &diContainer{}
}

func (d *diContainer) PaymentAPIV1() paymentV1.PaymentServiceServer {
	if d.paymentAPIV1 == nil {
		d.paymentAPIV1 = paymentAPIV1.NewAPI(d.Service())
	}
	return d.paymentAPIV1
}

func (d *diContainer) Service() service.PaymentService {
	if d.paymentService == nil {
		d.paymentService = paymentService.NewService(d.Repository())
	}
	return d.paymentService
}

func (d *diContainer) Repository() repository.PaymentRepository {
	if d.paymentRepository == nil {
		d.paymentRepository = payment.NewRepository()
	}
	return d.paymentRepository
}
