func lengthOfLIS(nums []int) int {
	if len(nums) < 0 {
		return 0
	}
	result := 1
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		result = max(result, dp[i])
	}
	return result
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}