package inventory

import (
	"context"

	"github.com/rocket-crm/inventory/internal/model"
	"github.com/rocket-crm/inventory/internal/repository/converter"
)

func(r *repository) ListParts(ctx context.Context, filter model.PartsFilter) ([]model.Part, error) {
	var result []model.Part
	
	if len(filter.Uuids) == 0 && len(filter.ManufacturerCountries) == 0 && len(filter.Names) == 0 && len(filter.Tags) == 0 && len(filter.Categories) == 0 {
		for _, p := range r.data {
			result = append(result, converter.PartToModel(p))
		}
		return result, nil
	}

	if len(filter.Uuids) > 0 {
		for _, v := range filter.Uuids {
			for _, p := range r.data {
				if p.Uuid == v {
					result = append(result, converter.PartToModel(p))
				}
			}
		}
	}

	if len(filter.Names) > 0 {
		for _, v := range filter.Names {
			for _, p := range r.data {
				if p.Name == v {
					result = append(result, converter.PartToModel(p))
				}
			}
		}
	}
	if len(filter.Categories) > 0 {
		for _, v := range filter.Categories {
			for _, p := range r.data {
				if v == model.Category(p.Category) {
					result = append(result, converter.PartToModel(p))
				}
			}
		}
	}
	if len(filter.ManufacturerCountries) > 0 {
		for _, v := range filter.ManufacturerCountries {
			for _, p := range r.data {
				if v == p.Manufacturer.Country {
					result = append(result, converter.PartToModel(p))
				}
			}
		}
	}
	if len(filter.Tags) > 0 {
		for _, v := range filter.Tags {
			for _, p := range r.data {
				for _, t := range p.Tags {
					if v == t {
						result = append(result, converter.PartToModel(p))
					}
				}
			}
		}
	}
	return result, nil
}