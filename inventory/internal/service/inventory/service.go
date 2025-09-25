package inventory

import "github.com/rocket-crm/inventory/internal/repository"




type service struct {
	inventoryRepository repository.InventoryRepository
}

func NewService(inventoryRepository repository.InventoryRepository) *service {
	return &service{inventoryRepository}
}