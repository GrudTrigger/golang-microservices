package v1

import (
	"context"

	inventoryV1 "github.com/rocker-crm/shared/pkg/proto/inventory/v1"
	"github.com/rocket-crm/inventory/internal/converter"
)

func(a *api) ListParts(ctx context.Context, req *inventoryV1.ListPartsRequest) (*inventoryV1.ListPartsResponse, error) {
	parts, err := a.inventoryService.ListParts(ctx, converter.PartsFilterToModel(req.Filter))
	if err != nil {
		return &inventoryV1.ListPartsResponse{}, err
	}
	return &inventoryV1.ListPartsResponse{Parts: converter.PartsToProto(parts)}, nil
}