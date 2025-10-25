package decoder

import (
	"github.com/rocker-crm/notifacation/internal/model"
	eventsV1 "github.com/rocker-crm/shared/pkg/proto/events/v1"
	"google.golang.org/protobuf/proto"
)

func Decode(data []byte) (model.OrderPaidEvent, error) {
	var pb eventsV1.OrderPaidRecorder

	if err := proto.Unmarshal(data, &pb); err != nil {
		return model.OrderPaidEvent{}, err
	}

	return model.OrderPaidEvent{
		EventUuid:       pb.EventUuid,
		OrderUuid:       pb.OrderUuid,
		UserUuid:        pb.UserUuid,
		PaymentMethod:   pb.PaymentMethod,
		TransactionUuid: pb.TransactionUuid,
	}, nil
}
