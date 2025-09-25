package v1

import (
	"context"
	"errors"

	inventoryV1 "github.com/rocker-crm/shared/pkg/proto/inventory/v1"
	"github.com/rocket-crm/inventory/internal/converter"
	"github.com/rocket-crm/inventory/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *api) GetPart(ctx context.Context, req *inventoryV1.GetPartRequest) (*inventoryV1.GetPartResponse, error) {
	part, err := a.inventoryService.GetPart(ctx, req.GetUuid())
	if err != nil {
		if errors.Is(err, model.ErrPartNotFound) {
			return nil, status.Errorf(codes.NotFound, "part with UUID %s not found", req.Uuid)
		}
		return nil, err
	}

	return &inventoryV1.GetPartResponse{Part: converter.PartToProto(part)}, nil
}