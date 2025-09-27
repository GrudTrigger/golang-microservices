package repository

import (
	"context"

	"github.com/rocket-crm/inventory/internal/model"
)

//go:generate ../../../bin/mockery

type InventoryRepository interface {
	ListParts(ctx context.Context, filter model.PartsFilter) ([]model.Part, error)
	GetPart(ctx context.Context, uuid string) (model.Part, error)
}
