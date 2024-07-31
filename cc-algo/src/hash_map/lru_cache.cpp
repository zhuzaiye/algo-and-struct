// Copyright (c) 2024/7/12 by hzzhu92
// LRU Cache

#include <unordered_map>
using namespace std;

// 双向链表节点
class Node {
public:
    int key, value;
    Node *prev{}, *next{};

    // 防止隐式类型转换
    // 参考: https://stackoverflow.com/questions/121162/what-does-the-explicit-keyword-mean
    // explicit用于修饰类的构造函数, 被修饰的构造函数的类
    // 不能发生相应的隐式类型转换, 只能显式方式进行类型转换
    // 1. explicit 只能用于类内部的构造函数声明上
    // 2. explicit 作用于单个参数的构造函数
    explicit Node(int k = 0, int v = 0) : key(k), value(v) {}
};

class LRUCache {
private:
    int capacity;
    Node *dummy;

    unordered_map<int, Node *> key_to_node;

    void remove(Node *x) {
        x->prev->next = x->next;
        x->next->prev = x->prev;
    }

    void push_front(Node *x) {
        x->prev = dummy;
        x->next = dummy->next;
        dummy->next->prev = x;
        dummy->next = x;
    }

    Node *get_node(int key) {
        auto it = key_to_node.find(key);
        if (it == key_to_node.end()) {
            return nullptr;
        }
        auto node = it->second;
        remove(node);
        push_front(node);
        return node;
    }
public:
    LRUCache(int capacity) : capacity(capacity), dummy(new Node()) {
        dummy->prev = dummy;
        dummy->next = dummy;
    }

    int get(int key) {
        auto node = get_node(key);
        return node ? node->value : -1;
    }

    void put(int key, int value) {
        auto node = get_node(key);
        if (node) {
            node->value = value;
            return;
        }

        key_to_node[key] = node = new Node(key, value);
        push_front(node);
        if (key_to_node.size() > capacity) {
            auto back_node = dummy->prev;
            key_to_node.erase(back_node->key);
            remove(back_node);  // 删除最近最少使用节点
            delete back_node; // 释放内存
        }
    }
};