package inventory

import (
	"context"

	"github.com/rocket-crm/inventory/internal/model"
	"github.com/rocket-crm/inventory/internal/repository/converter"
)

func (r *repository) GetPart(ctx context.Context, uuid string) (model.Part, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, p := range r.data {
		if p.Uuid == uuid {
			return converter.PartToModel(p), nil
		}
	}
	return model.Part{}, model.ErrPartNotFound
}
