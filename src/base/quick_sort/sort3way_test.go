package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort3way(t *testing.T) {
	result := []int{3, 4, 5, 1, 2, 3}
	Sort3way(result, 0, len(result)-1)
	assert.Equal(t, result, []int{1, 2, 3, 3, 4, 5})
}
