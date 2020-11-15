package minArray

func minArray(numbers []int) int {
	left := 0
	right := len(numbers) - 1
	for right > left {
		mid := (left + right) / 2
		if numbers[right] > numbers[mid] {
			right = mid
		} else if numbers[right] < numbers[mid] {
			left = mid + 1
		} else {
			right--
		}
	}
	return numbers[left]
}
