package test
import ("github.com/stretchr/testify/assert"
"testing"
)
func TestQuickSort(t *testing.T){
	nums :=  []int{2,3,1}
	assert.Equal(t,Quick_Sort(nums,0,2),[]int{1,2,3})
}