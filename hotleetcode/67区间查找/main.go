package main 
import (
	"sort"
	"fmt"
)	
func searchRange(nums []int, target int) []int {
    leftmost := sort.SearchInts(nums, target)
    if leftmost == len(nums) || nums[leftmost] != target {
        return []int{-1, -1}
    }
    rightmost := sort.SearchInts(nums, target + 1) - 1
    return []int{leftmost, rightmost}
}
func main(){
	nums := []int{2,3,5,6}
	fmt.Println(searchRange(nums,5))
}

