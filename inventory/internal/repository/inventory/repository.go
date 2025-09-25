package inventory

import (
	"sync"

	repoModel "github.com/rocket-crm/inventory/internal/repository/model"
)

type repository struct {
	mu    sync.RWMutex
	data map[string]repoModel.Part
}

func NewRepository() *repository {
	return &repository{
		data: make(map[string]repoModel.Part),
	}
}

// ------- Тестовые данные пока нет бд -------
//
// parts: map[string]*inventoryV1.Part{
// 	"part-1": {
// 		Uuid:          uuid.NewString(),
// 		Name:          "Engine Core",
// 		Description:   "High-performance spaceship engine core",
// 		Price:         1200.50,
// 		StockQuantity: 5,
// 		Category:      inventoryV1.Category_CATEGORY_ENGINE,
// 		Dimensions: &inventoryV1.Dimensions{
// 			Length: 2.5,
// 			Width:  1.2,
// 			Height: 1.8,
// 			Weight: 350.0,
// 		},
// 		Manufacturer: &inventoryV1.Manufacturer{
// 			Name:    "SpaceTech Inc",
// 			Country: "USA",
// 			Website: "https://spacetech.example.com",
// 		},
// 		Tags: []string{"engine", "spaceship", "core"},
// 		Metadata: map[string]*inventoryV1.Value{
// 			"serial": {Kind: &inventoryV1.Value_StringValue{StringValue: "SN-12345"}},
// 			"batch":  {Kind: &inventoryV1.Value_Int64Value{Int64Value: 42}},
// 		},
// 		CreatedAt: timestamppb.Now(),
// 		UpdatedAt: timestamppb.Now(),
// 	},
// 	"part-2": {
// 		Uuid:          uuid.NewString(),
// 		Name:          "Fuel Pump",
// 		Description:   "Reliable fuel pump for interstellar flights",
// 		Price:         300.75,
// 		StockQuantity: 12,
// 		Category:      inventoryV1.Category_CATEGORY_FUEL,
// 		Dimensions: &inventoryV1.Dimensions{
// 			Length: 0.8,
// 			Width:  0.4,
// 			Height: 0.6,
// 			Weight: 25.0,
// 		},
// 		Manufacturer: &inventoryV1.Manufacturer{
// 			Name:    "Galaxy Supplies",
// 			Country: "Germany",
// 			Website: "https://galaxy-supplies.example.com",
// 		},
// 		Tags: []string{"fuel", "pump"},
// 		Metadata: map[string]*inventoryV1.Value{
// 			"serial": {Kind: &inventoryV1.Value_StringValue{StringValue: "FP-00987"}},
// 		},
// 		CreatedAt: timestamppb.Now(),
// 		UpdatedAt: timestamppb.Now(),
// 	},
// 	"part-3": {
// 		Uuid:          uuid.NewString(),
// 		Name:          "Porthole Glass",
// 		Description:   "Reinforced glass for spaceship portholes",
// 		Price:         150.00,
// 		StockQuantity: 50,
// 		Category:      inventoryV1.Category_CATEGORY_PORTHOLE,
// 		Dimensions: &inventoryV1.Dimensions{
// 			Length: 1.0,
// 			Width:  1.0,
// 			Height: 0.02,
// 			Weight: 10.0,
// 		},
// 		Manufacturer: &inventoryV1.Manufacturer{
// 			Name:    "CosmoGlass",
// 			Country: "Japan",
// 			Website: "https://cosmoglass.example.com",
// 		},
// 		Tags: []string{"glass", "window"},
// 		Metadata: map[string]*inventoryV1.Value{
// 			"heat_resistance": {Kind: &inventoryV1.Value_DoubleValue{DoubleValue: 1250.5}},
// 		},
// 		CreatedAt: timestamppb.Now(),
// 		UpdatedAt: timestamppb.Now(),
// 	},
// },