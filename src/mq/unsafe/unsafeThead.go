package main
import (
	"fmt"
	"sync"
	//"time"
)
var wg sync.WaitGroup
func appendValue(array *[]int,i int){
	*array = append(*array,i)
	wg.Done()
}
func UnSafePrint(){
	var array []int
	for i :=0;i<100;i++{
		wg.Add(1)
		go appendValue(&array,i)
		
	}
	//time.Sleep(time.Second)
	wg.Wait()
	fmt.Println(array)
}
func main()  {
	UnSafePrint()
}

