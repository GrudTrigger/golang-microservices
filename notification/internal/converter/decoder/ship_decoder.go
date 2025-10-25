package decoder

import (
	"github.com/rocker-crm/notifacation/internal/model"
	eventsV1 "github.com/rocker-crm/shared/pkg/proto/events/v1"
	"google.golang.org/protobuf/proto"
)

func DecodeShipAssembled(data []byte) (model.ShipAssembledEvent, error) {
	var pb eventsV1.ShipAssembledRecorder

	if err := proto.Unmarshal(data, &pb); err != nil {
		return model.ShipAssembledEvent{}, err
	}

	return model.ShipAssembledEvent{
		EventUuid: pb.EventUuid,
		OrderUuid: pb.OrderUuid,
		UserUuid:  pb.UserUuid,
		BuildTime: pb.BuildTimeSec,
	}, nil
}
