func getIntersectionNode(headA, headB *ListNode) *ListNode {
	l1 := 1
	l2 := 1
	tmp1 := headA
	tmp2 := headB
	for tmp1 != nil {
		tmp1 = tmp1.Next
		l1 += 1
	}

	for tmp2 != nil {
		tmp2 = tmp2.Next
		l2 += 1
	}

	if l2 > l1 {
		offset := l2 - l1
		for offset > 0 {
			headB = headB.Next
			offset -= 1
		}
	} else if l2 < l1 {
		offset := l1 - l2
		for offset > 0 {
			headA = headA.Next
			offset -= 1
		}
	}

	// 一起走
	for headA != nil {

		if headA == headB {
			return headA
		} else {
			headA = headA.Next
			headB = headB.Next
		}

	}

	return nil
}