package strategy

import (
	"cmp"

	"github.com/GrayHat12/goche"
	"github.com/GrayHat12/goche/libs"
)

type FIFOStrategy[K cmp.Ordered, T any] struct {
	store *libs.LinkedList[K, T]
}

func NewFifoStrategy[K cmp.Ordered, T any](_ *goche.Cache[K, T]) goche.StrategyInterface[K, T] {
	return &FIFOStrategy[K, T]{
		store: libs.NewLinkedList[K, T](),
	}
}

func (s *FIFOStrategy[K, T]) Set(c *goche.Cache[K, T], id K, val T) {
	if s.store.Size() >= c.Max {
		s.store.RemoveFirst()
	}
	s.store.Add(id, val)
}

func (s *FIFOStrategy[K, T]) Get(c *goche.Cache[K, T], id K) (T, bool) {
	val, ok := s.store.Get(id)
	return val, ok
}

func (s *FIFOStrategy[K, T]) Remove(c *goche.Cache[K, T], id K) {
	s.store.Remove(id)
}
