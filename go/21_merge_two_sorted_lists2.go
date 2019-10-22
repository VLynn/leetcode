/**
 * 非递归解法，每次循环找出两个链表头更小的那个节点，作为结果链表的下一节点
 * 
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    headNode := ListNode{0, nil}
    index := &headNode
    
    // 每次循环找出两个链表头更小的那个节点，作为结果链表的下一节点
    // 某一链表遍历结束时，循环停止
    for l1 != nil && l2 != nil {
        if l1.Val <= l2.Val {
            index.Next = l1
            l1 = l1.Next
        } else {
            index.Next = l2
            l2 = l2.Next
        }
        index = index.Next
    }

    if l1 == nil {
        index.Next = l2
    }
    if l2 == nil {
        index.Next = l1
    }
    return headNode.Next
}