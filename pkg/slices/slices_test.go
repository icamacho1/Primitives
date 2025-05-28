package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {
	s := New("Tokyo", "Kyoto", "Osaka", "Hiroshima", "Nagasaki")

	assert.Equal(t, 5, s.Len())
	assert.Equal(t, "Osaka", s.Get(2))
	assert.Equal(t, "Tokyo", s.First())
	assert.Equal(t, "Nagasaki", s.Last())

	s.Append("Tokyo")
	assert.Equal(t, 6, s.Len())

	set, ok := s.Uniq()
	assert.True(t, ok)
	assert.Equal(t, 5, set.Len())

	s = s.Pop(2)
	assert.Equal(t, 5, s.Len())
	assert.Equal(t, "Hiroshima", s.Get(2))

	scopy := s.Map(func(item string, i int) string {
		return item + " Prefecture"
	})
	assert.Equal(t, "Tokyo Prefecture", scopy.Get(0))

}

func TestUnique(t *testing.T) {
	// Uniq hardcore test:
	mapSlice := New[map[string]int]()
	mapSlice.Append(map[string]int{
		"Tokyo":     1,
		"Nagasaki":  2,
		"Hiroshima": 3})
	mapSlice.Append(map[string]int{
		"Hiroshima": 3,
		"Nagasaki":  2,
		"Tokyo":     1,
	})
	for range 10_000 {
		uniqMapValues, ok := mapSlice.Uniq()
		assert.True(t, ok)
		assert.Equal(t, 1, uniqMapValues.Len())
		assert.Equal(t, map[string]int{
			"Nagasaki":  2,
			"Tokyo":     1,
			"Hiroshima": 3}, uniqMapValues.Get(0))
	}

	sliceSlice := New([]string{"Tokyo", "Nagasaki"}, []string{"Tokyo", "Nagasaki"})
	uniqSliceValues, ok := sliceSlice.Uniq()
	assert.True(t, ok)
	assert.Equal(t, 1, uniqSliceValues.Len())
	assert.Equal(t, []string{"Tokyo", "Nagasaki"}, uniqSliceValues.Get(0))
}

func TestForEach(t *testing.T) {
	intSlice := New(1, 2, 3)

	accum := 0
	intSlice.ForEach(func(item, _ int) {
		accum += item
	})

	assert.Equal(t, 6, accum)
}
