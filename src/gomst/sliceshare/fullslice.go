package main

import (
	"bytes"
	"fmt"
)

func main() {
	path := []byte("AAAA/BBBBBBBBB")
	spliteIndex := bytes.IndexByte(path, '/')
	str1 := path[:spliteIndex]
	str3 := path[:spliteIndex:spliteIndex]
	str2 := path[spliteIndex+1:]
	fmt.Printf("str1 = %v, Pointer = %p, len = %d, cap = %d\n", str1, &str1, len(str1), cap(str1))
	fmt.Printf("str3 = %v, Pointer = %p, len = %d, cap = %d\n", str3, &str3, len(str3), cap(str3))
	fmt.Println("str1 %v", string(str1))
	fmt.Println("str2 %v", string(str2))
	str1 = append(str1, "heihei"...)
	fmt.Println("str1 %v", string(str1))
	fmt.Println("str2 %v", string(str2))

}
