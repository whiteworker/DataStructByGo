package main
import (
	"fmt"
	"sync"
	//"time"
	"sort"
)
var wg sync.WaitGroup
func appendValue(array *[]int,i int){
	defer wg.Done()
	*array = append(*array,i)
	
}
func UnSafePrint(){
	var array []int
	wg.Add(100)
	for i :=0;i<100;i++{
		go appendValue(&array,i)
	}
	//time.Sleep(time.Second)
	wg.Wait()
	sort.Ints(array)
	fmt.Println(array)
}
func main()  {
	UnSafePrint()
}

