package goche

import (
	"errors"
	"sync"
)

type GenericFunction[T any] func() T

type Cache[T any] struct {
	store map[string]T
	mux   sync.RWMutex
	keys  []string
	index int
	max   int
}

func NewCache[T any](max int) *Cache[T] {
	return &Cache[T]{make(map[string]T), sync.RWMutex{}, make([]string, max), 0, max}
}

func (c *Cache[T]) removeKeyOnCurrentIndex() {
	key_on_current_index := c.keys[c.index]
	delete(c.store, key_on_current_index)
}

func (c *Cache[T]) Set(id string, value T) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.removeKeyOnCurrentIndex()
	c.keys[c.index] = id
	c.index = (c.index + 1) % c.max
	c.store[id] = value
}

func (c *Cache[T]) Get(id string) (T, error) {
	var none T
	c.mux.RLock()
	v, ok := c.store[id]
	c.mux.RUnlock()

	if !ok {
		return none, errors.New("a value with given key not found")
	}

	return v, nil
}
