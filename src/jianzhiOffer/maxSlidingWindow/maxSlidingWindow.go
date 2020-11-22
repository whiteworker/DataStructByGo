package main

import (
	"fmt"
)

func maxSlidingWindow(nums []int, k int) []int {
	if(len(nums)==0||k==0){
		return []int{}
	}
	windows,res := []int{},make([]int,len(nums)-k+1)
	for right,left := 0,1-k;right<len(nums);left++{
		//不在窗口
		if(left>0&&windows[0]==nums[left-1]){
			windows=windows[1:]
		}
		for(len(windows)>0&&windows[len(windows)-1]<nums[right]){
			windows=windows[:len(windows)-1]
		}
		windows=append(windows,nums[right])
		if(left>=0){
			res[left]=windows[0]
		}
		right++
	}
	return res

} 
func main(){
	nums :=[]int{1,3,-1,-3,5,3,6,7}
	fmt.Println(maxSlidingWindow(nums,3))
}

