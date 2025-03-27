package goche

import (
	"cmp"
	"errors"
	"sync"
)

type NewStrategyGenerator[K cmp.Ordered, T any] func(*Cache[K, T]) StrategyInterface[K, T]

type StrategyInterface[K cmp.Ordered, T any] interface {
	Set(*Cache[K, T], K, T)
	Get(*Cache[K, T], K) (T, bool)
	Remove(*Cache[K, T], K)
}

type Cache[K cmp.Ordered, T any] struct {
	mux      sync.RWMutex
	Max      int
	strategy StrategyInterface[K, T]
}

func NewCache[K cmp.Ordered, T any](max int, newStrategy NewStrategyGenerator[K, T]) *Cache[K, T] {
	var strategy StrategyInterface[K, T]
	cache := &Cache[K, T]{sync.RWMutex{}, max, strategy}
	strategy = newStrategy(cache)
	cache.strategy = strategy
	return cache
}

func (c *Cache[K, T]) Set(id K, value T) {
	c.mux.Lock()
	c.strategy.Set(c, id, value)
	defer c.mux.Unlock()
}

func (c *Cache[K, T]) Get(id K) (T, error) {
	var none T
	c.mux.RLock()
	v, ok := c.strategy.Get(c, id)
	c.mux.RUnlock()

	if !ok {
		return none, errors.New("a value with given key not found")
	}

	return v, nil
}

func (c *Cache[K, T]) Remove(id K) {
	c.mux.RLock()
	c.strategy.Remove(c, id)
	c.mux.RUnlock()
}

type Callable[P any, R any] func(P) R
type HashFunction[P any, K cmp.Ordered] func(P) K

type CachedFunction[K cmp.Ordered, P any, R any] struct {
	Call  Callable[P, R]
	Cache *Cache[K, R]
}

func FunctionDecorator[P any, R any, K cmp.Ordered](callable Callable[P, R], size int, strategyGenerator NewStrategyGenerator[K, R], hashFunction HashFunction[P, K]) CachedFunction[K, P, R] {
	var cache = NewCache(size, strategyGenerator)

	return CachedFunction[K, P, R]{
		Call: func(args P) R {
			// Hashing logic
			hash := hashFunction(args)

			// try getting result from cache
			response, err := cache.Get(hash)

			// if response in cache, return cached value
			if err == nil {
				return response
			}

			// else, fetch the live response
			response = callable(args)

			// store response in cache
			cache.Set(hash, response)

			return response
		},
		Cache: cache,
	}
}
