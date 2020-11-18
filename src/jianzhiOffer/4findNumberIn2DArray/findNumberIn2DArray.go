package findNumberIn2DArray

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	x, y := 0, len(matrix)-1
	for x < len(matrix[0]) && y >= 0 {
		if matrix[y][x] > target {
			y--
		} else if matrix[y][x] < target {
			x++
		} else {
			return true
		}
	}
	return false
}
