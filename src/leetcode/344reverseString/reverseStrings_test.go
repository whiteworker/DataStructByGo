package reverseString

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T){
    word :=[]byte("abcdefg")
	ReverseString(word)
	assert.Equal(t,word,[]byte("gfedcba"))
}

