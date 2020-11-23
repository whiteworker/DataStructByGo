package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHeadSort(t *testing.T){
	array := []int{
		55, 94,87,1,4,32,11,77,39,42,64,53,70,12,9,
	}
	HeapSort(array)
	assert.Equal(t,array,[]int{1 ,4,9 ,11 ,12, 32 ,39, 42, 53 ,55 ,64, 70 ,77, 87, 94})
}