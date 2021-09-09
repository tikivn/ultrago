package u_hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomString(t *testing.T) {
	t.Run("TestRandomString", func(t *testing.T) {
		text1 := RandomString(64)
		assert.Equal(t, 64, len(text1))
	})
}
