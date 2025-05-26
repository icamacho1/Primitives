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

	s2 := s.ForEach(func(item string, i int) string {
		return item + " City"
	})
	assert.Equal(t, "Tokyo City", s2.Get(0))
	assert.Equal(t, "Tokyo", s.Get(0))

	s.Map(func(item string, i int) string {
		return item + " Prefecture"
	})
	assert.Equal(t, "Tokyo Prefecture", s.Get(0))
}
