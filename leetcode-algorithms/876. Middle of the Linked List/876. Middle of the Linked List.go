func middleNode(head *ListNode) *ListNode {
    front := head

    for front != nil && front.Next != nil {
        front = front.Next.Next
        head = head.Next
    }

    return head

}