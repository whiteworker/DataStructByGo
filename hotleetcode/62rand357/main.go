package main
import ("fmt"
"math/rand"
)
func main(){
	 for i:=0;i<100;i++{
		//fmt.Println(rand3())
		 fmt.Println(rand5())
	 }
}
func rand3()int{
	return rand.Intn(3)
}
func rand5()int{
	for {
		 res :=(rand3()+1)*4+rand3()//（rand3()+1）*4 4 8 12    1~15
		 return (res+1)%5+1							
	}	
	
}