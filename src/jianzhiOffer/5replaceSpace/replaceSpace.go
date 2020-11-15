package replaceSpace

func replaceSpace(s string) string {
	var sb []rune
	for _, k := range s {
		if k == ' ' {
			sb = append(sb, '%', '2', '0')
		} else {
			sb = append(sb, k)
		}
	}
	return string(s)
}
