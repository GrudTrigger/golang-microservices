package v1

import (
	inventoryV1 "github.com/rocker-crm/shared/pkg/proto/inventory/v1"
	"github.com/rocket-crm/inventory/internal/service"
)



type api struct {
	inventoryV1.UnimplementedInventoryServiceServer
	inventoryService service.InventoryService
}


func NewAPI(inventoryService service.InventoryService) *api {
	return &api{inventoryService: inventoryService}
}