package twoNumbers
import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddTwoNumbers(t *testing.T){
	result := twoSum([]int{2,7,11,15},13)
	result1 := twoSum([]int{2,7,11,15},9)
	assert.Equal(t,result,[]int{0,2})
	assert.Equal(t,result1,[]int{0,1})
}