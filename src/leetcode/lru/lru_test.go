package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLRU(t *testing.T){
	lru := Constructor(2)
	lru.Put(1,1)
	lru.Put(2,2)
	lru.Put(3,3)
	assert.Equal(t,lru.Get(2),2)
	lru.Put(4,4)
	assert.Equal(t,lru.Get(3),-1)
	assert.Equal(t,lru.Get(2),2)
}