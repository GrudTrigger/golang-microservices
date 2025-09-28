package v1

import inventoryV1 "github.com/rocker-crm/shared/pkg/proto/inventory/v1"

type client struct {
	generatedClient inventoryV1.InventoryServiceClient
}

func NewClient(generatedClient inventoryV1.InventoryServiceClient) *client {
	return &client{generatedClient: generatedClient}
}
