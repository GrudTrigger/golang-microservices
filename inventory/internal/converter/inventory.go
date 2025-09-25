package converter

import (
	"log"

	inventoryV1 "github.com/rocker-crm/shared/pkg/proto/inventory/v1"
	"github.com/rocket-crm/inventory/internal/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func PartToProto(p model.Part) *inventoryV1.Part {
	var createdAt *timestamppb.Timestamp
	if p.CreatedAt != nil {
		createdAt = timestamppb.New(*p.CreatedAt)
	}

	var updatedAt *timestamppb.Timestamp
	if p.UpdatedAt != nil {
		updatedAt = timestamppb.New(*p.UpdatedAt)
	}

	return &inventoryV1.Part{
		Uuid: p.Uuid,
		Name: p.Name,
		Description: p.Description,
		Price: p.Price,
		StockQuantity: p.StockQuantity,
		Category: inventoryV1.Category(p.Category),
		Dimensions: DimensionsToProto(p.Dimensions),
		Manufacturer: ManufacturerToProto(p.Manufacturer),
		Tags: p.Tags,
		Metadata: MetadataToProto(p.Metadata),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}


func DimensionsToProto(d *model.Dimensions) *inventoryV1.Dimensions {
	return &inventoryV1.Dimensions{
		Length: d.Length,
		Width: d.Weight,
		Height: d.Height,
		Weight: d.Weight,
	}
}

func ManufacturerToProto(m *model.Manufacturer) *inventoryV1.Manufacturer {
	return &inventoryV1.Manufacturer{
		Name: m.Name,
		Country: m.Country,
		Website: m.Website,
	}
}

func MetadataToProto(m map[string]*any) map[string]*inventoryV1.Value {
	
	res := make(map[string]*inventoryV1.Value)

	for k, v := range m {
		switch t := (*v).(type) {
		case string:
			res[k] = &inventoryV1.Value{
				Kind: &inventoryV1.Value_StringValue{StringValue: t},
			}
		case int64:
			res[k] = &inventoryV1.Value{
				Kind: &inventoryV1.Value_Int64Value{Int64Value: t},
			}
		case float64:
			res[k] = &inventoryV1.Value{
				Kind: &inventoryV1.Value_DoubleValue{DoubleValue: t},
			}
		case bool:
			res[k] = &inventoryV1.Value{
				Kind: &inventoryV1.Value_BoolValue{BoolValue: t},
			}
		default:
			log.Printf("unsupported metadata type for key %s: %T", k, t)
		}
	}
	return res
}


func PartsFilterToModel(f *inventoryV1.PartsFilter) model.PartsFilter {
	return model.PartsFilter{
		Uuids: f.Uuids,
		Names: f.Names,
		Categories: CategoriesToModel(f.Categories),
		ManufacturerCountries: f.ManufacturerCountries,
		Tags: f.Tags,
	}
}

func CategoriesToModel(c []inventoryV1.Category) []model.Category {
	var res []model.Category
	for _, v := range c {
		res = append(res, model.Category(v))
	}
	return res
}

func PartsToProto(parts []model.Part) []*inventoryV1.Part {
	res := make([]*inventoryV1.Part, len(parts))
	for _, v := range parts {
		p := PartToProto(v)
		res = append(res, p)
	}
	return res
}