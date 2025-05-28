package slices

import (
	"crypto/sha256"
	"encoding/json"
	"reflect"
	"slices"

	"github.com/icamacho1/Primitives/pkg/maps"
)

func New[T any](items ...T) Slice[T] {
	s := make(Slice[T], len(items))
	copy(s, items)
	return s
}

type Slice[T any] []T

func (s Slice[T]) Len() int {
	return len(s)
}

func (s Slice[T]) First() T {
	return s[0]
}

func (s Slice[T]) Last() T {
	return s[len(s)-1]
}

func (s Slice[T]) Get(i int) T {
	return s[i]
}

func (s *Slice[T]) Append(item ...T) Slice[T] {
	*s = append(*s, item...)
	return *s
}

func (s *Slice[T]) Pop(i int) Slice[T] {
	*s = slices.Delete(*s, i, i+1)
	return *s
}

func (s Slice[T]) Has(item T) bool {
	for _, v := range s {
		if reflect.DeepEqual(v, item) {
			return true
		}
	}
	return false
}

func (s Slice[T]) ForEach(doer func(T, int) T) {
	sCopy := make(Slice[T], len(s))
	for i, item := range s {
		sCopy[i] = doer(item, i)
	}
}

func (s Slice[T]) Map(doer func(T, int) T) Slice[T] {
	for i, item := range s {
		s[i] = doer(item, i)
	}
	return s
}

func (s Slice[T]) Uniq() (Slice[T], bool) {
	lookup := maps.New[string, bool]()
	uniqSlice := make(Slice[T], 0, len(s))
	for _, item := range s {
		bytes, err := json.Marshal(item)
		if err != nil {
			return s, false
		}
		shaBytes := sha256.Sum256(bytes)
		hash := string(shaBytes[:])
		if !lookup.Has(hash) {
			uniqSlice = append(uniqSlice, item)
		}
		lookup.Add(string(hash[:]), true)
	}
	return uniqSlice, true
}
