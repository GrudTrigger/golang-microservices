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