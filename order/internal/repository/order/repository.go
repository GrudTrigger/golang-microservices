package order

import (
	"sync"

	repoModal "github.com/rocket-crm/order/internal/repository/model"
)

type repository struct {
	mu     sync.RWMutex
	orders map[string]repoModal.Order
}

func NewRepository() *repository {
	return &repository{
		orders: make(map[string]repoModal.Order),
	}
}
