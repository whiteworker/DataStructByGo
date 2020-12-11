package stringpratice

func Reverse(str string) string {
	res := []rune(str)
	if len(str) > 5000 || len(str) == 0 {
		return ""
	} else {
		l := len(str)
		for i := 0; i < l/2; i++ {
			res[i], res[l-i-1] = res[l-i-1], res[i]
		}
	}
	return string(res)
}
