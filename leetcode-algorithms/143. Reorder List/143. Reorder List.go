func revert(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    var pre *ListNode
    for head != nil {
        tmp := head.Next
        head.Next = pre
        pre = head
        head = tmp
    }

    return pre
}

func findMid(head *ListNode) *ListNode {
    front := head
    for front != nil && front.Next != nil {
        front = front.Next.Next
        head = head.Next   
    }
    return head
}

func reorderList(head *ListNode)  {
    if head == nil {
        return
    }
    mid := findMid(head)
    l1 := head
    l2 := revert(mid.Next)
    mid.Next = nil

    var tmp1,tmp2 *ListNode
    for l1 != nil && l2 != nil {
        tmp1 = l1.Next
        tmp2 = l2.Next

        l1.Next = l2
        l1 = tmp1

        l2.Next = l1
        l2 = tmp2

    }

}