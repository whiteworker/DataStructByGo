package main
import (
	"fmt"
	"sync"	
	"sort"
)
var wg sync.WaitGroup
var lock sync.Mutex
func appendValue(array *[]int,i int){
	defer wg.Done()
	lock.Lock()
	{
		*array = append(*array,i)
	}
	lock.Unlock()
	
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

