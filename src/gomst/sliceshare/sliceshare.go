package main

import (
	"fmt"
)

func main() {
	hello := []byte{'h', 'e', 'l', 'l', 'o'}
	helloGo := append(hello, ' ', 'g', 'o')
	helloGolang := append(helloGo, 'l', 'a', 'n', 'g')
	helloGoWorld := append(helloGo, ' ', 'w', 'o', 'l', 'd')
	fmt.Printf("%s\n", hello)
	fmt.Printf("%s\n", helloGolang)
	fmt.Printf("%s\n", helloGoWorld)
}
