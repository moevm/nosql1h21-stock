package repository

import (
	"nosql1h21-stock-backend/backend/internal/model"
	"sync"
)

type ValidTickerCache struct {
	data sync.Map
}

func NewCache() *ValidTickerCache {
	return &ValidTickerCache{data: sync.Map{}}
}

func (c *ValidTickerCache) Load(key string) ([]model.ValidTicker, bool) {
	value, ok := c.data.Load(key)
	if !ok {
		return []model.ValidTicker{}, false
	}
	p, ok := value.([]model.ValidTicker)
	return p, ok
}

func (c *ValidTickerCache) Store(key string, value []model.ValidTicker) {
	c.data.Store(key, value)
}
