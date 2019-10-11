/**
 * 遍历链表，将链表分拆成奇链表和偶链表，最后奇链表尾指向偶链表头
 * 空间复杂度O(1)，时间复杂度O(N)
 * 
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    
    oddHead, oddTail := head, head
    evenHead, evenTail := head.Next, head.Next
    
    // evenTail为nil，说明链表总数为奇数且已遍历完
    // evenTail.Next为nil，说明链表总数为偶数且已遍历完
    for evenTail != nil && evenTail.Next != nil {
        oddTail.Next = evenTail.Next
        oddTail = evenTail.Next
        evenTail.Next = oddTail.Next
        evenTail = oddTail.Next
    }
    oddTail.Next = evenHead
    
    return oddHead
}