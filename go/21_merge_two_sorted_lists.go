/**
 * 递归处理，每次递归找出链表合并后的头节点
 * 时间复杂度O(L1+L2)，L1是链表l1的长度
 * 
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    
    // 合并后的头节点是两个链表中较小的头节点
    var headNode *ListNode
    if l1.Val <= l2.Val {
        headNode = l1
        l1 = l1.Next
    } else {
        headNode = l2
        l2 = l2.Next
    }
    // 递归处理
    headNode.Next = mergeTwoLists(l1, l2)
    return headNode
}