package service

import (
	"context"

	"github.com/rocket-crm/inventory/internal/model"
)

type InventoryService interface {
	ListParts(ctx context.Context, filter model.PartsFilter) ([]model.Part, error)
	GetPart(ctx context.Context, uuid string) (model.Part, error)
}
