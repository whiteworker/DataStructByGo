package main

//在lt之前的(lo~lt-1)都小于中间值
//在gt之前的(gt+1~hi)都大于中间值
//在lt~i-1的都等于中间值
//在i~gt的都还不确定（最终i会大于gt，即不确定的将不复存在）
func Sort3way(s []int, lo, hi int) {
	if lo >= hi {
		return
	}
	v, lt, i, gt := s[lo], lo, lo+1, hi
	for i <= gt {
		if s[i] < v {
			swap(s, i, lt)
			lt++
			i++
		} else if s[i] > v {
			swap(s, i, gt)
			gt--
		} else {
			i++
		}
	}
	Sort3way(s, lo, lt-1)
	Sort3way(s, gt+1, hi)
}

func swap(s []int, i int, j int) {
	s[i], s[j] = s[j], s[i]
}
