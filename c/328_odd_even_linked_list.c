/**
 * 遍历链表，将链表分拆成奇链表和偶链表，最后奇链表尾指向偶链表头
 * 空间复杂度O(1)，时间复杂度O(N)
 *
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode* oddEvenList(struct ListNode* head){
    if (head == NULL) return NULL;
    struct ListNode *oddHead, *oddTail, *evenHead, *evenTail;
    oddHead = oddTail = head;
    evenHead = evenTail = head->next;
    
    while (evenTail != NULL && evenTail->next != NULL) {
        oddTail->next = evenTail->next;
        oddTail = evenTail->next;
        evenTail->next = oddTail->next;
        evenTail = oddTail->next;
    }
    oddTail->next = evenHead;
    
    return oddHead;
}