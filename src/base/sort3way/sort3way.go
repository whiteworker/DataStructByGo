package sort3way

//在lt之前的(start~lt-1)都小于中间值
//在gt之前的(gt+1~end)都大于中间值
//在lt~i-1的都等于中间值
//在i~gt的都还不确定（最终i会大于gt，即不确定的将不复存在）
func Sort3way(s []int, start, end int) {
	if start >= end {
		return
	}
	v, lt, i, gt := s[start], start, start+1, end
	for i <= gt {
		if s[i] < v {
			s[i], s[lt] = s[lt], s[i]
			lt++
			i++
		} else if s[i] > v {
			s[i], s[gt] = s[gt], s[i]
			gt--
		} else {
			i++
		}
	}
	Sort3way(s, start, lt-1)
	Sort3way(s, gt+1, end)
}
