/*

HashMap结构的存储本体就是一个数组。
于是建立Entry数组作为存储空间， 然后根据传入的key计算HashCode
作为数组的索引存入数据，读取的时候计算HashCode直接取值。

size保存当前的存储的键值对数量。
listSize保存当前数组的大小

*/

#ifndef _MAP_H
#define _MAP_H

typedef struct HashNode {
    void* key;              // key
    void* value;            // value
    struct HashNode* next;  // 冲突链表
} HashNode;

// 采用链地址法
typedef struct {
    int size;            // hash map 不重复no的数量
    HashNode** hashArr;  // 二维数组，存key值不重复的node，
                         // key重复的node链接在HashNode->next
} HashMap;

HashMap* CreateHashMap(int n);
int InsertHashMap(HashMap* hashMap, char* key, char* value);
char* GetHashMap(HashMap* hashMap, char* key);
void DeleteHashMap(HashMap* hashMap);
int RemoveHashMap(HashMap* hashMap, char* key);
void PrintHashMap(HashMap* hsahMap);

#endif