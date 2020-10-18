func lenOfNode(head *ListNode) (length int) {

    for head != nil {
        length++
        head = head.Next;
    }
    return length
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    length := lenOfNode(head)
    dummy := &ListNode{0,head}

    cur := dummy
    for i := 0; i < length-n; i++ {
        cur = cur.Next
    }
    cur.Next = cur.Next.Next
    return dummy.Next
}