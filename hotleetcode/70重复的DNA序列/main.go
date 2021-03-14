func findRepeatedDnaSequences(s string) []string {
    ans, d := []string{}, map[string]bool{}
    for i, n := 0, len(s) - 9; i < n; i++ {
        if _, ok := d[s[i: i + 10]]; ok {
            d[s[i: i + 10]] = true
        } else {
            d[s[i: i + 10]] = false
        }
    }
    for si, ok := range d {
        if ok {
            ans = append(ans, si)
        }
    }
    return ans
}
