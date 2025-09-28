package payment

import "github.com/rocket-crm/payment/internal/repository"

type service struct {
	paymentRepository repository.PaymentRepository
}

func NewService(paymentRepository repository.PaymentRepository) *service {
	return &service{paymentRepository}
}
