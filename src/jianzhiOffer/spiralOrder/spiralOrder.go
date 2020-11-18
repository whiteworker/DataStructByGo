package main

import (
	"fmt"
)


func spiralOrder(matrix [][]int) []int {
    if(len(matrix)==0){
        return []int{}
    } else{
        left,right,top,bottom,x :=0,len(matrix[0])-1,0,len(matrix)-1,0
        res := make([]int, (right+1)*(bottom+1))
        for{
            for i := left; i<right+1; i++ {
				res[x]=matrix[top][i]
				x++
			}
			top++
            if(top>bottom){
                break;
			}
			
            for i := top ;i<bottom+1;i++{
				res[x]=matrix[i][right]
				x++
			}
			right--
            if(left>right){
                break;
			}
			
            for i:= right;left<i+1;i--{
				res[x]=matrix[bottom][i]
				x++
			}
			bottom--
            if(top>bottom){
                break;
			}
			
            for i :=bottom;top<i+1;i--{
				res[x]=matrix[i][left]
				x++
			}
			left++
            if(left>right){
                break
			}
			
			fmt.Println(res)
		}
		return res
	}
}

func main(){
	res := spiralOrder([][]int{{1,2,3},{4,5,6},{7,8,9}})
	fmt.Println(res)
}
	
