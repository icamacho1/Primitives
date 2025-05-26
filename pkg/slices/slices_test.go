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
}
