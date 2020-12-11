package stringpratice

import (
	"strings"
)

func IsUnique(str string) bool {
	if len(str) > 3000 {
		return false
	}
	for _, v := range str {
		if v > 127 {
			return false
		}
		if strings.Count(str, string(v)) > 1 {
			return false
		}
	}
	return true

}
