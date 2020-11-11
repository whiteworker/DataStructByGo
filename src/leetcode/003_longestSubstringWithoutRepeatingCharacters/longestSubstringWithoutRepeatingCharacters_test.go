package longestSubstringWithoutRepeatingCharacters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
func TestLong(t *testing.T){
	result1 := lengthOfLongestSubstring("abcabcbb")
	result2 := lengthOfLongestSubstring("pwwkew")
	assert.Equal(t,result1,3)
	assert.Equal(t,result2,3)
}
