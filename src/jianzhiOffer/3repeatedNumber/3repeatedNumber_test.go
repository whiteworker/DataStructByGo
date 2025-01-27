package jianzhiOffer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T){
	result1 := findRepeatNumber1([]int{2,3,3,4,1})
	assert.Equal(t,result1,3)
	result2 := findRepeatNumber2([]int{2,3,3,4,1})
	assert.Equal(t,result2,3)
	result3 := findRepeatNumber3([]int{2,3,3,4,1})
	assert.Equal(t,result3,3)
}
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
			
            if(top>bottom){
                break;
			}
			top++
            for i := top ;i<bottom+1;i++{
				res[x]=matrix[i][right]
				x++
			}
			
            if(left>right){
                break;
			}
			right--
            for i:= right;left<i+1;i--{
				res[x]=matrix[bottom][i]
				x++
			}
			
            if(top>bottom){
                break;
			}
			bottom--
            for i :=bottom;top<i+1;i--{
				res[x]=matrix[i][left]
				x++
			}
			
            if(left>right){
                break
			}
			left++
			fmt.Printf("res:%v",res)
		}
		return res
	}
}