package maps

func New[K comparable, V any]() Map[K, V] {
	return make(Map[K, V])
}

type Map[K comparable, V any] map[K]V

func (m Map[K, V]) Len() int {
	return len(m)
}

func (m Map[K, V]) Get(key K) (V, bool) {
	value, ok := m[key]
	return value, ok
}

func (m Map[K, V]) Has(key K) bool {
	_, ok := m[key]
	return ok
}

func (m Map[K, V]) Add(key K, value V) Map[K, V] {
	m[key] = value
	return m
}

func (m Map[K, V]) MustGet(key K) V {
	return m[key]
}

func (m Map[K, V]) Keys() []K {
	keys := make([]K, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

func (m Map[K, V]) Pop(key K) Map[K, V] {
	delete(m, key)
	return m
}
