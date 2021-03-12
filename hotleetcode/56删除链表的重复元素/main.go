
	func deleteDuplicates(head *ListNode) *ListNode {
		if head == nil || head.Next == nil {
			return head
		}

		head.Next = deleteDuplicates(head.Next)
		if head.Val == head.Next.Val {
			head = head.Next
		}

		return head

	}