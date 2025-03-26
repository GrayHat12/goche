package strategy

import "github.com/GrayHat12/goche"

type FIFOStrategy[T any] struct {
	store map[string]T
	keys  []string
	index int
}

func NewFifoStrategy[T any](cache *goche.Cache[T]) goche.StrategyInterface[T] {
	return &FIFOStrategy[T]{
		store: make(map[string]T),
		keys:  make([]string, cache.Max),
		index: 0,
	}
}

func (s *FIFOStrategy[T]) removeKeyOnCurrentIndex() {
	key_on_current_index := s.keys[s.index]
	delete(s.store, key_on_current_index)
}

func (s *FIFOStrategy[T]) Set(c *goche.Cache[T], id string, val T) {
	s.removeKeyOnCurrentIndex()
	s.keys[s.index] = id
	s.index = (s.index + 1) % c.Max
	s.store[id] = val
}

func (s *FIFOStrategy[T]) Get(c *goche.Cache[T], id string) (T, bool) {
	val, ok := s.store[id]
	return val, ok
}
