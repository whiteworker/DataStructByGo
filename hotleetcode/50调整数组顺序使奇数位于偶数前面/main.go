func exchange(nums []int) []int {
	left := 0
	right := len(nums) - 1
	for left < right {
		for left < right && nums[left]%2 == 1 {
			left++
		}
		for left < right && nums[right]%2 == 0 {
			right--
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	return nums
}