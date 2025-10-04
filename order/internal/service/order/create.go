package order

import (
	"context"
	"errors"

	"github.com/rocket-crm/order/internal/model"
)

func (s *service) CreateOrder(ctx context.Context, data model.CreateOrder) (model.ResponseCreateOrder, error) {
	resParts, err := s.inventoryClient.ListParts(ctx, model.PartsFilter{Uuids: data.PartUuids})
	if err != nil {
		return model.ResponseCreateOrder{}, err
	}

	if len(resParts) != len(data.PartUuids) {
		return model.ResponseCreateOrder{}, errors.New("найдены не все запчасти, проверьте uuid деталей")
	}
	var totalPrice float32

	for _, p := range resParts {
		totalPrice += p.Price
	}

	o, err := s.orderRepository.Create(ctx, data, totalPrice)
	if err != nil {
		return model.ResponseCreateOrder{}, err
	}

	return model.ResponseCreateOrder{UUID: o.OrderUUID, TotalPrice: totalPrice}, nil
}
