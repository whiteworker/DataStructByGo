func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil {
		return false
	}
	xlen := len(matrix)
	if xlen == 0 {
		return false
	}
	//二分查找
	ylen := len(matrix[0])
	x, y := xlen-1, 0
	for x >= 0 && y <= ylen-1 {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] > target {
			x--
		} else {
			y++
		}
	}
	return false
}

