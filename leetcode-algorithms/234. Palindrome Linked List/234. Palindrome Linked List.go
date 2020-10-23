func isPalindrome(head *ListNode) bool {
	vals := make([]int, 0)

	for head != nil {
		vals = append(vals, head.Val)
		head = head.Next
	}

	n := len(vals)
	for i, v := range vals[:n/2] {
		if v != vals[n-1-i] {
			return false
		}
	}
	return true
}