package shellsort

func Shellsort(array []int) {
	l := len(array)
	h := 1
	for h < l/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < l; i++ {
			for j := i; j >= h; j -= h {
				if array[j-h] > array[j] {
					array[j-h], array[j] = array[j], array[j-h]
				} else {
					break
				}
			}
		}
		h = h / 3
	}

}
