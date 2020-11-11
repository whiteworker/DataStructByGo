package main
import(
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)
func TestQuickSort(t *testing.T){
	list := []int{5}
	fmt.Printf("org:%v\n",list)
	Quick_Sort(list,0,len(list)-1)
	fmt.Printf("result:%v\n",list)
	assert.Equal(t,list,[]int{5})
	list2 := []int{2,5,2,3,5,7,1}
	fmt.Printf("org:%v\n",list2)
	Quick_Sort(list2,0,len(list2)-1)
	fmt.Printf("result:%v\n",list2)
	assert.Equal(t,list2,[]int{1,2,2,3,5,5,7})
}