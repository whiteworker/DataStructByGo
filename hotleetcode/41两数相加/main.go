
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2, sum := 0, 0, 0
		if l1 != nil {
			n1 = l1.Val
		}
		if l2 != nil {
			n2 = l2.Val
		}
		sum = n1 + n2 + carry
		carry = sum / 10
		if head == nil {
			head = &ListNode{Val: sum % 10}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum % 10}
			tail = tail.Next
		}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
		tail = tail.Next
	}
	return head
}