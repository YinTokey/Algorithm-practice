func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    var ptr1 *ListNode = &ListNode{Val:0,Next:nil}
    head := ptr1

    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            ptr1.Next = l1
            l1 = l1.Next
        } else {
            ptr1.Next = l2
            l2 = l2.Next
        }
        ptr1 = ptr1.Next
    }

    if l1 != nil {
        ptr1.Next = l1
    }

    if l2 != nil {
        ptr1.Next = l2
    }

    return head.Next
}