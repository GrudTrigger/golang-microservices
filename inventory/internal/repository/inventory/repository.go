package inventory

import (
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/rocket-crm/inventory/internal/repository/converter"
	repoModel "github.com/rocket-crm/inventory/internal/repository/model"
)

type repository struct {
	mu   sync.RWMutex
	data map[string]repoModel.Part
}

func NewRepository() *repository {
	return &repository{
		data: map[string]repoModel.Part{
			"part-1": {
				Uuid:          uuid.NewString(),
				Name:          "Engine Core",
				Description:   "High-performance spaceship engine core",
				Price:         1200.50,
				StockQuantity: 5,
				Category:      1,
				Dimensions: &repoModel.Dimensions{
					Length: 2.5,
					Width:  1.2,
					Height: 1.8,
					Weight: 350.0,
				},
				Manufacturer: &repoModel.Manufacturer{
					Name:    "SpaceTech Inc",
					Country: "USA",
					Website: "https://spacetech.example.com",
				},
				Tags: []string{"engine", "spaceship", "core"},
				Metadata: map[string]any{
					"serial": "SN-12345",
					"batch":  int64(42),
				},
				CreatedAt: converter.Ptr(time.Now()),
				UpdatedAt: converter.Ptr(time.Now()),
			},
			"part-2": {
				Uuid:          uuid.NewString(),
				Name:          "Fuel Pump",
				Description:   "Reliable fuel pump for interstellar flights",
				Price:         300.75,
				StockQuantity: 12,
				Category:      2,
				Dimensions: &repoModel.Dimensions{
					Length: 0.8,
					Width:  0.4,
					Height: 0.6,
					Weight: 25.0,
				},
				Manufacturer: &repoModel.Manufacturer{
					Name:    "Galaxy Supplies",
					Country: "Germany",
					Website: "https://galaxy-supplies.example.com",
				},
				Tags: []string{"fuel", "pump"},
				Metadata: map[string]any{
					"serial": "FP-00987",
				},
				CreatedAt: converter.Ptr(time.Now()),
				UpdatedAt: converter.Ptr(time.Now()),
			},
			"part-3": {
				Uuid:          uuid.NewString(),
				Name:          "Porthole Glass",
				Description:   "Reinforced glass for spaceship portholes",
				Price:         150.00,
				StockQuantity: 50,
				Category:      3,
				Dimensions: &repoModel.Dimensions{
					Length: 1.0,
					Width:  1.0,
					Height: 0.02,
					Weight: 10.0,
				},
				Manufacturer: &repoModel.Manufacturer{
					Name:    "CosmoGlass",
					Country: "Japan",
					Website: "https://cosmoglass.example.com",
				},
				Tags: []string{"glass", "window"},
				Metadata: map[string]any{
					"heat_resistance": 1250.5,
				},
				CreatedAt: converter.Ptr(time.Now()),
				UpdatedAt: converter.Ptr(time.Now()),
			},
		},
	}
}
