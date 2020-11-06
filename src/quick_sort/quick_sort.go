package main
import "fmt"
func Quick_Sort(array []int, start int, end int){
	if (start<end) {
		i :=start
		j :=end
		key :=array[start]
		for(i<j){
			//右边小
			for(i<j&&array[j]>=key){
				j--
			}
			if (i<j) {
				array[i]=array[j]
				i++
			}
			for(i<j&&array[i]<=key){
				i++
			}
			if(i<j){
				array[j]=array[i]
				j--
			}
		}
		array[i]=key;
		Quick_Sort(array,start,i-1)
		Quick_Sort(array,i+1,end)
	}
}
func main(){
	var array=[]int{5,1,6,3,7,2,7,23,2}
	Quick_Sort(array,0,len(array)-1)
	fmt.Println(array);
}