package goche

import (
	"errors"
	"sync"
)

type NewStrategyGenerator[T any] func(*Cache[T]) StrategyInterface[T]

type StrategyInterface[T any] interface {
	Set(*Cache[T], string, T)
	Get(*Cache[T], string) (T, bool)
}

type Cache[T any] struct {
	mux      sync.RWMutex
	Max      int
	strategy StrategyInterface[T]
}

func NewCache[T any](max int, newStrategy NewStrategyGenerator[T]) *Cache[T] {
	var strategy StrategyInterface[T]
	cache := &Cache[T]{sync.RWMutex{}, max, strategy}
	strategy = newStrategy(cache)
	cache.strategy = strategy
	return cache
}

func (c *Cache[T]) Set(id string, value T) {
	c.mux.Lock()
	c.strategy.Set(c, id, value)
	defer c.mux.Unlock()
}

func (c *Cache[T]) Get(id string) (T, error) {
	var none T
	c.mux.RLock()
	v, ok := c.strategy.Get(c, id)
	c.mux.RUnlock()

	if !ok {
		return none, errors.New("a value with given key not found")
	}

	return v, nil
}
