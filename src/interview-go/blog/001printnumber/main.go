package main

import (
	"fmt"
	"sync"
)

//使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母， 最终效果如下：
//12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
func main() {
	letter, number := make(chan bool), make(chan bool)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		i := 1
		for {
			_, ok := <-number
			if ok {
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
			} else {
				fmt.Print("???")
			}
		}
	}()
	go func(wg *sync.WaitGroup) {
		str := "ABCDEFGHIJKLMNOPQRSTUVWSYZ"
		i := 0
		for {
			_, ok := <-letter
			if ok {
				if i > len(str)-1 {
					wg.Done()
					return
				}
				fmt.Print(str[i : i+1])
				i++
				if i > len(str)-1 {
					i = 0
				}
				fmt.Print(str[i : i+1])
				i++
				number <- true
			} else {
				fmt.Print("???")
			}
		}

	}(&wg)
	number <- true
	wg.Wait()
}
