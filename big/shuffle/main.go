package main 
import (
	"fmt"
	"time"
	"math/rand"
)
func main(){
	fmt.Println(shuffle([]int{1,2,3,4,5,6}))
	fmt.Println(shuffle2([]int{1,2,3,4,5,6}))
}
func shuffle(nums []int) []int {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := len(nums); i > 0; i-- {
		last := i - 1
		idx := rand.Intn(i)
		nums[last], nums[idx] = nums[idx], nums[last]
	}
	return nums
}
func shuffle2(nums []int) []int {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := len(nums); i > 0; i-- {
		last := i - 1
		idx := rand.Perm(i)
		nums[last], nums[idx[i-1]] = nums[idx[i-1]], nums[last]
	}
	return nums
}

	