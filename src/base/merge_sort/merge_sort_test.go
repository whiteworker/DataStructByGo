package merge_sort
import(
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T){
	list2 := []int{2,5,2,3,5,7,1}
	fmt.Printf("org:%v\n",list2)
	list2=Merge_Sort(list2)
	fmt.Printf("result:%v\n",list2)
	assert.Equal(t,list2,[]int{1,2,2,3,5,5,7})
}

