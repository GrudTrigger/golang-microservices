package inventory

import (
	"context"

	"github.com/rocket-crm/inventory/internal/model"
)

func (s *service) ListParts(ctx context.Context, filter model.PartsFilter) ([]model.Part, error) {
	parts, err := s.inventoryRepository.ListParts(ctx, filter)
	if err != nil {
		return []model.Part{}, err
	}
	return parts, err
}
