func getDecimalValue(head *ListNode) int {
    // 1. 翻转链表
    if head == nil || head.Next == nil {
        return head.Val
    }

    var prev *ListNode
    cur := head
    for cur != nil {
        cur.Next = prev
        prev = cur
        cur = cur.Next
    }
    //  pre 为翻转后的头结点
    
    // 2. 按照公式换算
    var result int = 0
    var count float64 = 0
    for prev != nil {
        result += prev.Val * int(math.Pow(2,count))
        prev = prev.Next
        count += 1.0
    }
    return result
}