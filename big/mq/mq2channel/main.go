package main

import (
	"fmt"
	"time"
)

func Producer(cdata chan<- int, step int) {
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
	cdata := make(chan int, 10)
	go Producer(cdata, 2)
	go Producer(cdata, 3)
	go Consumer(cdata)
	time.Sleep(time.Second * 5)
}
