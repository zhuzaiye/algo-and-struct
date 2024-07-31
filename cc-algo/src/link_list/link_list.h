// Copyright (c) 2024/7/30 by hzzhu92
//

#ifndef LINK_LIST_H_
#define LINK_LIST_H_

struct ListNode {
    int val;         // 节点值
    ListNode *next;  // 指向下一个节点
    explicit ListNode(int x) : val(x), next(nullptr) {}  // 构造函数
};

#endif //LINK_LIST_H_
