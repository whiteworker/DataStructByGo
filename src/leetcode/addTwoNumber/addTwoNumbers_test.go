package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumber(t *testing.T) {
	l1 := &ListNode{Value: 2, Next: &ListNode{Value: 4, Next: &ListNode{Value: 3}}}
	l2 := &ListNode{Value: 5, Next: &ListNode{Value: 6, Next: &ListNode{Value: 4}}}
	result := AddTwoNumber(l1, l2)
	assert.Equal(t, result, &ListNode{Value: 7, Next: &ListNode{Value: 0, Next: &ListNode{Value: 8}}})
}
