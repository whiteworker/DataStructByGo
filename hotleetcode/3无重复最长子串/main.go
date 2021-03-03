package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	length := len(s)
	res, rk := 0, -1
	for i := 0; i < length; i++ {
		if i > 0 {
			delete(m, s[i-1])
		}
		for rk+1 < length && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
		}
		res = max(res, rk-i+1)
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	str := "abbcdeabbc"
	fmt.Println(lengthOfLongestSubstring(str))
}
