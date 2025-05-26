package slices

import "slices"

func has[T any, C comparable](s Slice[T], v T) bool {
	sCopy := make(Slice[C], len(s))
	for i, item := range s {
		v, ok := any(item).(C)
		if !ok {
			return false
		}

		sCopy[i] = v
	}

	return slices.Contains(sCopy, any(v).(C))
}

func uniq[T any, C comparable](s Slice[T]) (Slice[T], bool) {
	lookup := make(map[C]bool)
	uniqSlice := make(Slice[T], 0, len(s))
	for _, item := range s {
		v, ok := any(item).(C)
		if !ok {
			return uniqSlice, false
		}

		if !lookup[v] {
			uniqSlice = append(uniqSlice, item)
			lookup[v] = true
		}
	}

	return uniqSlice, true
}
