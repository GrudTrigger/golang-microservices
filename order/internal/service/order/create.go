package order

import (
	"context"
	"errors"

	inventoryV1 "github.com/rocker-crm/shared/pkg/proto/inventory/v1"
	"github.com/rocket-crm/order/internal/model"
)

func (s *service) CreateOrder(ctx context.Context, data model.CreateOrder) (model.ResponseCreateOrder, error) {
	resParts, err := s.inventoryClient.ListParts(ctx, &inventoryV1.ListPartsRequest{Filter: &inventoryV1.PartsFilter{Uuids: data.PartUuids}})
	if err != nil {
		return model.ResponseCreateOrder{}, err
	}

	if len(resParts.Parts) != len(data.PartUuids) {
		return model.ResponseCreateOrder{}, errors.New("найдены не все запчасти, проверьте uuid деталей")
	}
	var totalPrice float32

	for _, p := range resParts.Parts {
		totalPrice += p.Price
	}

	o := s.orderRepository.Create(data, totalPrice)

	return model.ResponseCreateOrder{UUID: o.OrderUUID, TotalPrice: totalPrice}, nil
}
