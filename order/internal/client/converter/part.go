package converter

import (
	"time"

	inventoryV1 "github.com/rocker-crm/shared/pkg/proto/inventory/v1"
	"github.com/rocket-crm/order/internal/model"
)

type Category int32

func CategoriesModelToProto(c []model.Category) []inventoryV1.Category {
	var res []inventoryV1.Category

	for _, v := range c {
		res = append(res, inventoryV1.Category(v))
	}
	return res
}

func PartsFilterToProto(filter model.PartsFilter) *inventoryV1.PartsFilter {
	return &inventoryV1.PartsFilter{
		Uuids:                 filter.Uuids,
		Names:                 filter.Names,
		Categories:            CategoriesModelToProto(filter.Categories),
		ManufacturerCountries: filter.ManufacturerCountries,
		Tags:                  filter.Tags,
	}
}

func ManufacturerToModel(m *inventoryV1.Manufacturer) *model.Manufacturer {
	return &model.Manufacturer{
		Name:    m.Name,
		Country: m.Country,
		Website: m.Website,
	}
}

func DimensionsToModel(d *inventoryV1.Dimensions) *model.Dimensions {
	return &model.Dimensions{
		Length: d.Length,
		Width:  d.Width,
		Weight: d.Weight,
		Height: d.Height,
	}
}

func MetadataToModel(m map[string]*inventoryV1.Value) map[string]any {
	result := make(map[string]any, len(m))

	for k, v := range m {
		if v == nil {
			result[k] = nil
			continue
		}

		switch val := v.Kind.(type) {
		case *inventoryV1.Value_StringValue:
			result[k] = val.StringValue
		case *inventoryV1.Value_Int64Value:
			result[k] = val.Int64Value
		case *inventoryV1.Value_DoubleValue:
			result[k] = val.DoubleValue
		case *inventoryV1.Value_BoolValue:
			result[k] = val.BoolValue
		default:
			result[k] = nil
		}
	}
	return result
}

func PartConvertToModel(p *inventoryV1.Part) model.Part {
	var createdAt *time.Time
	if p.CreatedAt != nil {
		t := p.CreatedAt.AsTime()
		createdAt = &t
	}

	var updatedAt *time.Time
	if p.UpdatedAt != nil {
		t := p.UpdatedAt.AsTime()
		updatedAt = &t
	}
	return model.Part{
		Uuid:          p.Uuid,
		Name:          p.Name,
		Description:   p.Description,
		Price:         p.Price,
		StockQuantity: p.StockQuantity,
		Category:      model.Category(p.Category),
		Tags:          p.Tags,
		Metadata:      MetadataToModel(p.Metadata),
		Manufacturer:  ManufacturerToModel(p.Manufacturer),
		Dimensions:    DimensionsToModel(p.Dimensions),
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}
}

func PartListToModel(parts []*inventoryV1.Part) []model.Part {
	var res []model.Part
	for _, p := range parts {
		res = append(res, PartConvertToModel(p))
	}
	return res
}
