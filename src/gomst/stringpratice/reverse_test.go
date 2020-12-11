package stringpratice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	str := "abcd"
	assert.Equal(t, Reverse(str), "dcba")
	str1 := "abcd1"
	assert.Equal(t, Reverse(str1), "1dcba")
}
