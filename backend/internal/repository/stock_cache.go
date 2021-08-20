package repository

import (
	"nosql1h21-stock-backend/backend/internal/model"
	"sync"
)

type Cache struct {
	data sync.Map
}

func NewCache() *Cache {
	return &Cache{data: sync.Map{}}
}

func (c *Cache) Load(key string) ([]model.ValidTicker, bool) {
	value, ok := c.data.Load(key)
	if !ok {
		return []model.ValidTicker{}, false
	}
	p, ok := value.([]model.ValidTicker)
	return p, ok
}

func (c *Cache) Store(key string, value []model.ValidTicker) {
	c.data.Store(key, value)
}
