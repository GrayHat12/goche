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
	Mux      sync.RWMutex
	Max      int
	Strategy StrategyInterface[T]
}

func NewCache[T any](max int, newStrategy NewStrategyGenerator[T]) *Cache[T] {
	var strategy StrategyInterface[T]
	cache := &Cache[T]{sync.RWMutex{}, max, strategy}
	strategy = newStrategy(cache)
	cache.Strategy = strategy
	return cache
}

func (c *Cache[T]) Set(id string, value T) {
	c.Mux.Lock()
	c.Strategy.Set(c, id, value)
	defer c.Mux.Unlock()
}

func (c *Cache[T]) Get(id string) (T, error) {
	var none T
	c.Mux.RLock()
	v, ok := c.Strategy.Get(c, id)
	c.Mux.RUnlock()

	if !ok {
		return none, errors.New("a value with given key not found")
	}

	return v, nil
}
