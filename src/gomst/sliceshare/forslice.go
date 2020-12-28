package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40}
	for _, value := range slice {
		value += 233
		//fmt.Printf("value = %d , value-addr = %x , slice-addr = %x\n", value, &value, &slice[index])
	}
	fmt.Printf("slice:%v", slice)
	for index, _ := range slice {
		v := &slice[index]
		*v = 233
		//fmt.Printf("value = %d , value-addr = %x , slice-addr = %x\n", value, &value, &slice[index])
	}
	fmt.Printf("slice:%v", slice)
}
