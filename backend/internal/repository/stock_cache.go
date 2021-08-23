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

func (c *ValidTickerCache) Load(key string) (model.ValidData, bool) {
	value, ok := c.data.Load(key)
	if !ok {
		return model.ValidData{}, false
	}
	p, ok := value.(model.ValidData)
	return p, ok
}

func (c *ValidTickerCache) Store(key string, value model.ValidData) {
	c.data.Store(key, value)
}
