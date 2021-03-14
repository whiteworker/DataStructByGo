package main

// 带缓冲区的channel

import (
	"fmt"
	"sync"
)

func produce(wg *sync.WaitGroup, ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
		fmt.Println("Send:", i)
	}
	wg.Done()
}

func consumer(wg *sync.WaitGroup, ch <-chan int) {
	for i := 0; i < 100; i++ {
		v := <-ch
		fmt.Println("Receive:", v)
	}
	wg.Done()
}
func Producer(cdata chan int, step int) {
	for i := 0; ; i += step {
		cdata <- i
	}
}
func Consumer(cdata chan int) {
	for {
		if v, ok := <-cdata; ok {
			fmt.Println(v)
		}
	}
}
func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 10)
	wg.Add(2)
	go produce(&wg, ch)
	go consumer(&wg, ch)
	wg.Wait()
	fmt.Println("ok")
}
