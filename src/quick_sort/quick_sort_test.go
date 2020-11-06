package main
import(
	"fmt"
	"reflect"
	"testing"
)
func TestQuickSort(t *testing.T){
	list := []int{5}
	fmt.Printf("org:%v\n",list)
	Quick_Sort(list,0,len(list)-1)
	fmt.Printf("result:%v\n",list)
	if reflect.DeepEqual(list,[]int{5}){
		fmt.Printf("QuickSort List1 Test list is Success\n")
	} else{
		t.Logf("QuickSort List1 Test is Fail\n")
	}
	list2 := []int{2,5,2,3,5,7,1}
	fmt.Printf("org:%v\n",list2)
	Quick_Sort(list2,0,len(list2)-1)
	fmt.Printf("result:%v\n",list2)
	if reflect.DeepEqual(list2,[]int{1,2,2,3,5,5,7}){
		fmt.Printf("QuickSort List2 Test is Success\n")
	} else{
		t.Logf("QuickSort List2 Test is Fail\n")
	}
}