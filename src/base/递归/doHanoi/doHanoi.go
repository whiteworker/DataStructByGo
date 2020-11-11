package main
import "fmt"
///(片的编号，源柱，目的柱，辅助柱)
///n :1  doHanoi(1,A,B,C)
///n :2  doHanoi(1,A,C,B) 2 doHanoi(1,C,B,A)
///n :3  doHanoi(2,A,C,B) 3 doHanoi(2,C,B,A) doHanoi()
///n :n  doHanoi(n-1,A,C,B) n A->B doHanoi(n-1,C,B,A)
func doHanoi(n int ,src ,dest,hpr string){
	if  n>1{
		doHanoi(n-1,src,hpr,dest)
		fmt.Printf("Move No . %d dish from %s to %s \n",n,src,dest)
		doHanoi(n-1,hpr,dest,src)
	} else if n==1{
		fmt.Printf("Move No. %d dish from %s to %s\n", n, src, dest)
	} 
}
func main(){
	doHanoi(3, "A", "B", "C")
}