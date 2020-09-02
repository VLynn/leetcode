/**
 * 因为是有序的，所以相同数值是聚在一起的，遍历一遍列表即可
 * 
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    current := head
    for current != nil && current.Next != nil {
        // 如果next和current值相同，删除next结点；否则移动current到下一个结点
        if current.Val == current.Next.Val {
            current.Next = current.Next.Next
        } else {
            current = current.Next
        }
    }
    return head
}