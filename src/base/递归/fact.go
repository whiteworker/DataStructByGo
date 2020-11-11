package main
import "fmt"
//阶乘
func fact(n int )(int){
	if n<=1{
		return n
	} else{
		return n*fact(n-1)
	}
}
func step(n int,stepTemp map[int]int) int {
	if n<=0{
		return 0
	}
	// 取缓存
	if value , ok := stepTemp[n]; ok{
		return value
	}
	switch n{
	case 1:
		stepTemp[1]=1
		return 1
	case 2:
		value := 1+step(1,stepTemp)
		stepTemp[2]=value
		return value
	case 3:
		value := 1+step(2,stepTemp)+step(1,stepTemp)
		stepTemp[3]=value
		return value
	default :
		value := step(n-3,stepTemp)+step(n-2,stepTemp)+step(n-1,stepTemp)
		stepTemp[n]=value
		return value
	}
}
func main(){
	fmt.Println(fact(5))
	// 动态规划，空间换时间
	stepTemp := make(map[int]int)
	fmt.Println(step(4,stepTemp))
	fmt.Println(step(5,stepTemp),stepTemp)
}