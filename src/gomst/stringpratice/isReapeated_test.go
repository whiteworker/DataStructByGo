package stringpratice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsUnique(t *testing.T) {
	assert.False(t, IsUnique("aba"))
	assert.True(t, IsUnique("abc"))
}
