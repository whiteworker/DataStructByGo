package jianzhiOffer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T){
	result1 := findRepeatNumber1([]int{2,3,3,4,1})
	assert.Equal(t,result1,3)
	result2 := findRepeatNumber2([]int{2,3,3,4,1})
	assert.Equal(t,result2,3)
	result3 := findRepeatNumber3([]int{2,3,3,4,1})
	assert.Equal(t,result3,3)
}