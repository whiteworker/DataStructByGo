package Print1toMax

func printNumbers(n int) []int {
	sum := 1
	for i := 0; i < n; i++ {
		sum *= 10
	}
	result := [sum - 1]int{}
	for i := 0; i < sum-1; i++ {
		result[i] = i + 1
	}
}
