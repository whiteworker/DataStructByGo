package main

import (
	"fmt"
	"testing"
)

func TestIface1(t *testing.T) {

}
func main() {
	var (
		data  *int
		eface interface{}
	)
	eface = data
	fmt.Println(data == nil)
	fmt.Println(eface == nil)
}
