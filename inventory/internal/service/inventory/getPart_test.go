package inventory

import (
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/rocket-crm/inventory/internal/model"
	"github.com/rocket-crm/inventory/internal/repository/converter"
)

func (s *ServiceSuite) TestGetPartSuccess() {
	partUuid := gofakeit.UUID()
	p := model.Part{
		Uuid:          partUuid,
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
	}

	s.inventoryRepository.On("GetPart", s.ctx, partUuid).Return(p, nil)
	part, err := s.service.GetPart(s.ctx, partUuid)
	s.Require().NoError(err)
	s.Require().Equal(partUuid, part.Uuid)
}

func (s *ServiceSuite) TestGetPartNotFound() {
	partUuid := gofakeit.UUID()
	s.inventoryRepository.On("GetPart", s.ctx, partUuid).Return(model.Part{}, model.ErrPartNotFound)
	part, err := s.service.GetPart(s.ctx, partUuid)
	s.Require().ErrorIs(err, model.ErrPartNotFound)
	s.Require().Equal(model.Part{}, part)
}
