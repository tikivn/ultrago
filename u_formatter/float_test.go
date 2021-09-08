package u_formatter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatFloat(t *testing.T) {
	x1 := FormatFloat(12423.45353, 0)
	x2 := FormatFloat(12423.75353, 0)
	assert.Equal(t, float64(12423), x1)
	assert.Equal(t, float64(12424), x2)
}
