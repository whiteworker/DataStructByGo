func longestPalindrome(s string) string {
	n := len(s)
	res := ""
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for L := 0; L < n; L++ {
		for i := 0; i+L < n; i++ {
			j := i + L
			if L == 0 {
				dp[i][j] = 1
			} else if L == 1 {
				if s[i] == s[j] {
					dp[i][j] = 1
				}

			} else {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] > 0 && L+1 > len(res) {
				res = s[i : i+L+1]
			}
		}
	}
	return res
}