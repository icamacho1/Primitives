package sets

import (
	"github.com/icamacho1/Primitives/pkg/maps"
	"github.com/icamacho1/Primitives/pkg/slices"
)

func NewOrdered[T comparable](items ...T) OrderedSet[T] {
	os := OrderedSet[T]{
		items:  slices.New[T](),
		lookup: maps.New[T, int](),
	}

	return *os.Append(items...)
}

type OrderedSet[T comparable] struct {
	items  slices.Slice[T]
	lookup maps.Map[T, int]
}

func (os OrderedSet[T]) Len() int {
	return len(os.items)
}

func (os *OrderedSet[T]) Append(items ...T) *OrderedSet[T] {
	idx := os.Len()
	for _, item := range items {
		if !os.lookup.Has(item) {
			os.items.Append(item)
			os.lookup.Add(item, idx)
			idx++
		}
	}
	return os
}

func (os OrderedSet[T]) Get(i int) T {
	return os.items.Get(i)
}

func (os OrderedSet[T]) Last() T {
	return os.items.Get(os.items.Len() - 1)
}

func (os OrderedSet[T]) First() T {
	return os.items.Get(0)
}

func (os OrderedSet[T]) Has(item T) bool {
	return os.lookup.Has(item)
}

func (os *OrderedSet[T]) Pop(item T) *OrderedSet[T] {
	if idx, hasValue := os.lookup.Get(item); hasValue {
		// Get the items that will be shifted first
		i := idx
		for _, item := range os.items[idx:] {
			os.lookup.Add(item, i-1)
			i++
		}
		os.lookup.Pop(item)
		os.items.Pop(idx)
	}
	return os
}
