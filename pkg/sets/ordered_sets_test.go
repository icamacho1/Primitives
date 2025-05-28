package sets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOrdered(t *testing.T) {
	items := []string{"Tokyo", "Osaka", "Tokyo"}
	os := NewOrdered(items...)

	assert.Equal(t, 2, os.Len())
	assert.Equal(t, "Osaka", os.Get(1))

	os.Append("Tokyo", "Kyoto", "Nagasaki")
	assert.Equal(t, 4, os.Len())

	os.Pop("Tokyo")
	assert.Equal(t, 3, os.Len())

	os.Append("Tokyo")
	assert.Equal(t, "Kyoto", os.Get(1))
	assert.Equal(t, "Tokyo", os.Last())
}
