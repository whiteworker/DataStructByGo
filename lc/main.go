package main
import "fmt"


func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
func test1(){
	res := make ([]int,5)
	res=append(res,1,2,3)
	fmt.Println(res)
}
func test2(){
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}
func main() {
	test2()
}