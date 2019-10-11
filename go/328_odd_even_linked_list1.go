/**
 * 遍历链表，奇数节点插入到已处理过的奇数节点后，偶数节点保持不动
 * 空间复杂度O(1)，时间复杂度O(N)
 * 
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    isOdd := true
    // 当前节点，当前节点的上一个节点，已处理过的最大奇数节点
    var curNode, lastNode, curOddNode *ListNode
    
    // 遍历链表，奇数节点插入到已处理过的奇数节点后，偶数节点保持不动
    for curNode = head; curNode != nil; {
        if isOdd {
            tmp := curNode
            if curOddNode != nil {
                lastNode.Next = curNode.Next
                curNode = curNode.Next
                tmp.Next = curOddNode.Next
                curOddNode.Next = tmp
            } else {
                // 首节点只移动前后指针
                lastNode = curNode
                curNode = curNode.Next
            }
            curOddNode = tmp
        } else {
            lastNode = curNode
            curNode = curNode.Next
        }
        
        isOdd = !isOdd
    }
    
    return head
}