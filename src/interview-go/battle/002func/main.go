package main

import (
	"fmt"
	"time"
)

func main() {
	values := []string{"a", "b", "c"}
	for _, v := range values {
		go func() {
			fmt.Println(v)
		}()
	}
	for _, v := range values {
		//顺序不确定 Go是没有引用传递的，即使是foo1()在参数上加了*，内部实现机制仍旧是值传递
		go func(v string) {
			fmt.Println(v)
		}(v)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("end")

}
