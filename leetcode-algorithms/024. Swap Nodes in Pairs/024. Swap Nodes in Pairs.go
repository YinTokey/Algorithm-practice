func swapPairs(head *ListNode) *ListNode {

	dummy := &ListNode{0, head}
	tmp := dummy
	for tmp.Next != nil && tmp.Next.Next != nil {
		ptr1 := tmp.Next
		ptr2 := tmp.Next.Next

		tmp.Next = ptr2
		ptr1.Next = ptr2.Next
		ptr2.Next = ptr1

		tmp = ptr1
	}
	return dummy.Next
}