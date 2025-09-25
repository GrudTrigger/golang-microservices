package payment

import (
	"context"
	"log"

	"github.com/rocket-crm/payment/internal/model"
)

func(s *service) PayOrder(ctx context.Context, payOrder model.PayOrder) (string, error) {
	transactionUuid, err := s.paymentRepository.PayOrder(ctx, payOrder)
	if err != nil {
		return "", err
	}
	log.Printf("Оплата прошла успешно, transaction_uuid: <%s>\n", transactionUuid)
	
	return transactionUuid, nil
}