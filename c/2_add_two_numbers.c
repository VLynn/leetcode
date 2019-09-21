/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */

// 新建链表节点
struct ListNode* newNode() {
    struct ListNode *node = (struct ListNode*) malloc(sizeof(struct ListNode));
    node->next = NULL;
    return node;
}

// 链表反转
struct ListNode* reverseList(struct ListNode* list) {
    struct ListNode *p, *q, *head;
    head = newNode();
    p = list;
    
    while (p != NULL) {
        q = p;
        p = p->next;
        q->next = head->next;
        head->next = q;
    }
    return head->next;
}

struct ListNode* addTwoNumbers(struct ListNode* l1, struct ListNode* l2){
    struct ListNode *res = newNode();
    // 当前位数值，进位数值
    int currentPosition, carryPosition;
    currentPosition = carryPosition = 0;
    
    // 从低位往高位遍历，直到两个数都遍历完
    while (l1 != NULL || l2 != NULL) {
        // 当前位非空，累加并前移一位
        if (l1 != NULL) {
            currentPosition += l1->val;
            l1 = l1->next;
        }
        if (l2 != NULL) {
            currentPosition += l2->val;
            l2 = l2->next;
        }
        
        // 当前位 + 上一位进位
        currentPosition += carryPosition;
        // 判断是否进位
        if (currentPosition > 9) {
            currentPosition -= 10;
            carryPosition = 1;
        } else {
            carryPosition = 0;
        }
        // 新增节点，并插入结果链表
        struct ListNode* node = newNode();
        node->val = currentPosition;
        node->next = res->next;
        res->next = node;
        // 重置当前位
        currentPosition = 0;
    }
    
    // 判断是否有最高位进位
    if (carryPosition != 0) {
        struct ListNode* node = newNode();
        node->val = carryPosition;
        node->next = res->next;
        res->next = node;
    }
    
    //反转链表
    return reverseList(res->next);
}