package inventory

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rocket-crm/inventory/internal/model"
	"github.com/rocket-crm/inventory/internal/repository/converter"
)

func (s *ServiceSuite) TestListPartsEmptyFiltersSuccess() {
	filter := model.PartsFilter{}
	parts := []model.Part{
		{
			Uuid:          uuid.NewString(),
			Name:          "Engine Core",
			Description:   "High-performance spaceship engine core",
			Price:         1200.50,
			StockQuantity: 5,
			Category:      1,
			Dimensions: &model.Dimensions{
				Length: 2.5,
				Width:  1.2,
				Height: 1.8,
				Weight: 350.0,
			},
			Manufacturer: &model.Manufacturer{
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
		{
			Uuid:          uuid.NewString(),
			Name:          "Warp Drive",
			Description:   "Advanced warp drive for interstellar travel",
			Price:         2500.75,
			StockQuantity: 3,
			Category:      2,
			Dimensions: &model.Dimensions{
				Length: 3.0,
				Width:  1.5,
				Height: 2.0,
				Weight: 420.0,
			},
			Manufacturer: &model.Manufacturer{
				Name:    "Galactic Motors",
				Country: "UK",
				Website: "https://galacticmotors.example.com",
			},
			Tags: []string{"warp", "engine", "spaceship"},
			Metadata: map[string]any{
				"serial": "SN-67890",
				"batch":  int64(84),
			},
			CreatedAt: converter.Ptr(time.Now()),
			UpdatedAt: converter.Ptr(time.Now()),
		},
	}

	s.inventoryRepository.On("ListParts", s.ctx, filter).Return(parts, nil)
	p, err := s.service.ListParts(s.ctx, filter)
	s.Require().Equal(p, parts)
	s.NoError(err)
}

var responsePart = model.Part{
	Uuid:          uuid.NewString(),
	Name:          "Warp Drive",
	Description:   "Advanced warp drive for interstellar travel",
	Price:         2500.75,
	StockQuantity: 3,
	Category:      2,
	Dimensions: &model.Dimensions{
		Length: 3.0,
		Width:  1.5,
		Height: 2.0,
		Weight: 420.0,
	},
	Manufacturer: &model.Manufacturer{
		Name:    "Galactic Motors",
		Country: "UK",
		Website: "https://galacticmotors.example.com",
	},
	Tags: []string{"warp", "engine", "spaceship"},
	Metadata: map[string]any{
		"serial": "SN-67890",
		"batch":  int64(84),
	},
	CreatedAt: converter.Ptr(time.Now()),
	UpdatedAt: converter.Ptr(time.Now()),
}

func (s *ServiceSuite) TestListPartsWithFilterSuccess() {

	tests := []struct {
		filter   model.PartsFilter
		resParts []model.Part
	}{
		{
			filter:   model.PartsFilter{Tags: []string{"warp"}},
			resParts: []model.Part{responsePart},
		},
		{
			filter:   model.PartsFilter{Names: []string{"Warp Drive"}},
			resParts: []model.Part{responsePart},
		},
		{
			filter:   model.PartsFilter{Uuids: []string{responsePart.Uuid}},
			resParts: []model.Part{responsePart},
		},
		{
			filter:   model.PartsFilter{ManufacturerCountries: []string{"UK"}},
			resParts: []model.Part{responsePart},
		},
	}

	for _, test := range tests {
		s.inventoryRepository.On("ListParts", s.ctx, test.filter).Return(test.resParts, nil)
		p, err := s.service.ListParts(s.ctx, test.filter)
		s.Require().Equal(p, test.resParts)
		s.Require().NoError(err)
	}
}

func (s *ServiceSuite) TestListPartsWithFilterEmptyResponseSuccess() {
	filter := model.PartsFilter{
		Names: []string{"empty response"},
	}
	s.inventoryRepository.On("ListParts", s.ctx, filter).Return([]model.Part{}, nil)
	p, err := s.service.ListParts(s.ctx, filter)
	s.Require().Len(p, 0)
	s.Require().NoError(err)
}

func (s *ServiceSuite) TestListPartsWithError() {
	filter := model.PartsFilter{
		Names: []string{"Warp Drive"},
	}
	s.inventoryRepository.On("ListParts", s.ctx, filter).Return([]model.Part{}, errors.New("test error"))
	p, err := s.service.ListParts(s.ctx, filter)
	s.Require().Error(err)
	s.Require().Len(p, 0)
}
