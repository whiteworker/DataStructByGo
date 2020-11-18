package shellsort

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShellsort(t *testing.T) {
	array := []int{5, 4, 3, 2, 1, 6, 7}
	Shellsort(array)
	fmt.Println(array)
	assert.Equal(t, array, []int{1, 2, 3, 4, 5, 6, 7})
}
