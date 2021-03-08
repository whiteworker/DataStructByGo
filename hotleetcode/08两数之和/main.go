package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hashtable := map[int]int{}
	for i, num := range nums {
		if value, ok := hashtable[target-num]; ok {
			return []int{i, value}
		} else {
			hashtable[num] = i
		}
	}
	return []int{}
}
func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(twoSum(nums, 3))
}
