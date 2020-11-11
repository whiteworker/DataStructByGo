package jianzhiOffer

import (
	"sort"
)

func findRepeatNumber1(nums []int)int{
	hashMap := make(map[int]int)
	for _,value :=range nums{
		if _,ok :=hashMap[value];ok{
			return value
		} else {
			hashMap[value]=1
		}
	}
	return -1
}
func findRepeatNumber2(nums []int)int{
	sort.Ints(nums)
	for i:=0;i<len(nums);i++{
		if(nums[i]==nums[i+1]){
			return nums[i]
		}
	}
	return -1
}
func findRepeatNumber3(nums []int)int{
	for i := 0; i < len(nums); i++ {
        for i != nums[i] {
            if nums[i] == nums[nums[i]] {
                return nums[i]
            }
            nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
        }
    }
    return -1

}