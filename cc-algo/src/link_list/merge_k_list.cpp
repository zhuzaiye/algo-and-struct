// Copyright (c) 2024/7/30 by hzzhu92
// 合并k个有序链表

#include <vector>
#include <queue>
#include "link_list.h"

using namespace std;

class Solution {
public:
    ListNode *mergeKLists(vector<ListNode *> &lists) {
        if (lists.empty()) {
            return nullptr;
        }
        return mergeBinary(lists, 0, lists.size() - 1);
    }
private:
    ListNode *mergeTwoList(ListNode *l1, ListNode *l2) {
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

        prev->next = l1 == nullptr ? l2 : l1;
        ListNode *r = preHead->next;
        delete preHead;
        return r;
    }

    // 递归是针对能明确量小的情况下实现
    ListNode *mergeBinary(vector<ListNode *> &lists, int l, int r) {
        if (l == r) {
            return lists[l];
        }
        if (l > r) {
            return nullptr;
        }
        int mid = (l + r) >> 1;
        return mergeTwoList(mergeBinary(lists, l, mid), mergeBinary(lists, mid + 1, r));
    }
};

class Solution2 {
public:
    struct status {
        int val;
        ListNode *ptr;
        // 重载比较运算符`<`, 用于比较两个`status`结构实例
        // `const status &rhs`中`rhs`是一个对`status`结构体的常量引用, 表示要与当前的实例进行比较的另一个`status`实例
        // `const`关键字放在成员函数未部, 表示这个成员函数不会修改当前实例的成员变量
        bool operator<(const status &rhs) const {
            // 实现从小到大的排序
            return val > rhs.val;
        }
    };

    // 优先级队列, 先进先出, 在队列的基础上增加定义的排序功能, 本质是一个堆实现
    // 在queue头文件中，能够在O(1)的时间内找到最大的元素（默认为大顶堆），并且能在在O(logn)的时间内插入和删除元素
    // 方法:
    //     top 访问对头元素
    //     empty 队列是否为空
    //     size 返回队列内元素个数
    //     push 插入元素到队尾并排序
    //     emplace 原地构造一个元素并插入队列
    //     pop 弹出对头元素
    //     swap 交换内容
    priority_queue<status> q;

    ListNode *mergeKLists(vector<ListNode *> &lists) {
        if (lists.empty()) {
            return nullptr;
        }

        for (auto node : lists) {
            if (node) {
                q.push({node->val, node});
            }
        }

        auto *head = new ListNode(-1);
        ListNode *tail = head;
        while (!q.empty()) {
            // 获取当前最小node
            auto f = q.top();
            // 弹出该最小node
            q.pop();
            tail->next = f.ptr;
            tail = tail->next;
            // 并将该队列后续的node入队进行重新排序, 便于后续的比较组合
            if (f.ptr->next) {
                q.push({f.ptr->next->val, f.ptr->next});
            }
        }
        return head->next;
    }
};
