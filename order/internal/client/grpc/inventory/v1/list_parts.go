package v1

import (
	"context"

	inventoryV1 "github.com/rocker-crm/shared/pkg/proto/inventory/v1"
	"github.com/rocket-crm/order/internal/client/converter"
	"github.com/rocket-crm/order/internal/model"
)

func (c *client) ListParts(ctx context.Context, filter model.PartsFilter) ([]model.Part, error) {
	p, err := c.generatedClient.ListParts(ctx, &inventoryV1.ListPartsRequest{
		Filter: converter.PartsFilterToProto(filter),
	})
	if err != nil {
		return nil, err
	}
	return converter.PartListToModel(p.Parts), nil
}
