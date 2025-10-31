package session

import (
	"github.com/rocker-crm/platform/pkg/cache"
)

const (
	cacheKeyPrefix = "session:"
)

type repository struct {
	cache cache.RedisClient
}

func NewRepository(cache cache.RedisClient) *repository {
	return &repository{
		cache: cache,
	}
}
