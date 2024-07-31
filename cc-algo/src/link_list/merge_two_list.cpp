// Copyright (c) 2024/7/29 by hzzhu92
// 合并两个有序链表为一个有序链表

#include "link_list.h"

class Solution {
public:
    ListNode *mergeTwoLists(ListNode *l1, ListNode *l2) {
        auto *preHead = new ListNode(-1);
        ListNode *prev = preHead;
        while (l1 != nullptr && l2 != nullptr) {
            if (l1->val < l2->val) {
                prev->next = l1;
                l1 = l1->next;
            } else {
                prev->next = l2;
                l2 = l2->next;
            }
            prev = prev->next;
        }

        // 合并长度较长的那一个
        prev->next = l1 == nullptr ? l2 : l1;

        auto r = preHead->next;
        // 释放哨兵占位节点内存,
        // 如果使用智能指针则不需要
        delete preHead;
        return r;
    }
};