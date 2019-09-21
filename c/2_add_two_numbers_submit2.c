/**
 * 使用较长的那个输入链表存储结果，代码可读性会差一些，但节省空间
 * 时间复杂度O(N)，空间复杂度O(1)
 * Runtime: 8 ms, faster than 96.92% of C online submissions for Add Two Numbers.
 * Memory Usage: 8.5 MB, less than 92.00% of C online submissions for Add Two Numbers.
 *
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
    struct ListNode *index, *head1, *head2;
    // 输入的链表长度
    int len1, len2;
    // 当前位数值，进位数值
    int currentPosition, carryPosition;
    
    head1 = l1;
    head2 = l2;
    len1 = len2 = 0;
    currentPosition = carryPosition = 0;
    
    // 从低位往高位遍历，直到两个数都遍历完
    while (l1 != NULL || l2 != NULL) {
        // 当前位非空，累加
        if (l1 != NULL)
            currentPosition += l1->val;
        if (l2 != NULL)
            currentPosition += l2->val;
        // 当前位 + 上一位进位
        currentPosition += carryPosition;
        // 判断是否进位
        if (currentPosition > 9) {
            currentPosition -= 10;
            carryPosition = 1;
        } else {
            carryPosition = 0;
        }

        // 更新节点值，并前移一位
        if (l1 != NULL) {
            l1 -> val = currentPosition;
            index = l1;
            l1 = l1->next;
            len1++;
        }
        if (l2 != NULL) {
            l2 -> val = currentPosition;
            index = l2;
            l2 = l2->next;
            len2 ++;
        }
        // 重置当前位
        currentPosition = 0;

        // 判断是否有最高位进位
        if (index->next == NULL && carryPosition != 0) {
            struct ListNode* node = newNode();
            node->val = carryPosition;
            index->next = node;
        }

        // 如果只在一个链表上遍历，且没有进位了，无需继续遍历下去
        if ((l1 == NULL || l2 == NULL) && carryPosition == 0) {
            if (l1 == NULL)
                len2++;
            if (l2 == NULL)
                len1++;
            break;
        }
    }

    // 长度更长的链表存储了结果，默认l2
    return len2 >= len1 ? head2 : head1;
}