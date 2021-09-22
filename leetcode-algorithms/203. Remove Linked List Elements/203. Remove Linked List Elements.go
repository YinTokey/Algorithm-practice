func removeElements(head *ListNode, val int) *ListNode {
	pre := &ListNode{
		Val:  0,
		Next: head,
	}
	cur := pre

	for cur.Next != nil {
		if cur.Next.Val != val {
			cur = cur.Next
		} else {
			cur.Next = cur.Next.Next
		}
	}

	return pre.Next
}