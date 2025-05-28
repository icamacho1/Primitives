package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOverride(t *testing.T) {
	m := New[string, []int]()
	m.Add("test", []int{1, 2, 3})

	assert.Equal(t, Map[string, []int]{"test": {1, 2, 3}}, m)

	m.Override("test", func(current []int) []int {
		return append(current, 2)
	})

	assert.Equal(t, Map[string, []int]{"test": {1, 2, 3, 2}}, m)
}
