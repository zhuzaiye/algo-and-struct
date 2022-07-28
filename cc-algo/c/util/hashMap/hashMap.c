#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "hashMap.h"
#include "hashUtil.h"

/* 创建一个node数量为n的hash表 */
HashMap* CreateHashMap(int n) {
    HashMap* hashMap = (HashMap*)calloc(1, sizeof(HashMap));
    hashMap->hashArr = (HashMap**)calloc(n, sizeof(HashMap*));

    if (hashMap == NULL || hashMap->hashArr == NULL) {
        return NULL;
    }
    hashMap->size = n;
    return hashMap;
};

/* 插入键值对
 * 如果key相同，则后插入的value会覆盖已有的value
 */
int InsertHashMap(HashMap* hashMap, char* key, char* value) {
    HashNode* node = (HashNode*)calloc(1, sizeof(HashNode));
    if (node == NULL) {
        return 0;
    }

    // 复制key和value
    node->key = strdup(key);
    node->value = strdup(value);
    node->next = NULL;

    //
    int index = hash(key) % hashMap->size;
    if (hashMap->hashArr[index] == NULL) {
        hashMap->hashArr[index] = node;
    } else {
        // 用于遍历node的临时游标
        HashNode* temp = hashMap->hashArr[index];
        // 记录上一个node
        HashNode* prev = temp;
        // 循环遍历至最后一个节点node_end的next
        while (temp != NULL) {
            // 如果两个key相同，则直接用新node的value覆盖旧的
            if (strcmp(temp->key, node->key) == 0) {
                // 释放旧value
                free(temp->value);
                // 复制新value
                temp->value = strdup(node->value);
                return 1;
            }
            prev = temp;
            temp = temp->next;
        }
        // 最后一个节点node_end的next指向新建的node
        prev->next = node;
    }
    return 1;
}

/* 通过key获取value */
char* GetHashMap(HashMap* hashMap, char* key) {
    // 对hash结果求余，获取key位置
    int index = hash(key) % hashMap->size;
    // 用于遍历node的临时游标
    HashNode* temp = hashMap->hashArr[index];
    // 循环遍历至最后一个节点node_end的next
    while (temp != NULL) {
        // 如果两个key相同，则用新node的value覆盖旧的
        if (strcmp(temp->key, key) == 0) {
            return temp->value;
        }
        temp = temp->next;
    }
    return NULL;
}

/* 释放hash map 内存 */
void DeleteHashMap(HashMap* hashMap) {
    for (int i = 0; i < hashMap->size; i++) {
        // 用于遍历node的临时游标
        HashNode* temp = hashMap->hashArr[i];
        HashNode* prev = temp;
        // 循环遍历至最后一个节点node_end的next
        while (temp != NULL) {
            prev = temp;
            temp = temp->next;
            myFree(prev->key);
            myFree(prev->value);
            myFree(prev);
        }
    }
    myFree(hashMap->hashArr);
    myFree(hashMap);
    hashMap->hashArr = NULL;
    hashMap = NULL;
}

int RemoveHashMap(HashMap* hashMap, char* key) {
    // 对hash结果求余，获取key位置
    int index = hash(key) % hashMap->size;
    // 用于遍历node的临时游标
    HashNode* temp = hashMap->hashArr[index];
    if (temp == NULL) {
        return 0;
    }

    // 如果第一个就匹配中
    if (strcmp(temp->key, key) == 0) {
        // printf("找到:%s->%s\n", temp->key, temp->value);
        hashMap->hashArr[index] = temp->next;
        free(temp->key);
        free(temp->value);
        free(temp);
        return 1;
    } else {
        HashNode* prev = temp;
        temp = temp->next;
        // 循环遍历至最后一个节点node_end的next
        while (temp != NULL) {
            // 如果两个key相同，则用新node的value覆盖旧的
            if (strcmp(temp->key, key) == 0) {
                // printf("找到:%s->%s\n", temp->key, temp->value);
                prev->next = temp->next;
                free(temp->key);
                free(temp->value);
                free(temp);
                return 1;
            }
            prev = temp;
            temp = temp->next;
        }
    }
    return 0;
}

void PrintHashMap(HashMap* hashMap)
{
    printf("===========PrintHashMap==========\n");
    HashNode* node = NULL;
    for (int i = 0; i < hashMap->size; i++)
    {
        node = hashMap->hashArr[i];
        if (node != NULL) {
            HashNode* temp = node;
            while (temp != NULL) {
                printf("[%d]: %s -> %s\t", i, temp->key, temp->value);
                temp = temp->next;
            }
            printf("\n");
        }
    }
    printf("===============End===============\n");
}

void hashMapTest(void)
{
    HashMap* hashMap = CreateHashMap(5);
    if (hashMap) {
        printf("创建完成\n");
    }
    
    InsertHashMap(hashMap, "a", "a1");
    InsertHashMap(hashMap, "a", "a2");
    InsertHashMap(hashMap, "b", "b1");
    InsertHashMap(hashMap, "b", "b2");
    InsertHashMap(hashMap, "c", "c1");
    InsertHashMap(hashMap, "d", "d1");
    InsertHashMap(hashMap, "e", "e1");
    InsertHashMap(hashMap, "f", "f1");
    InsertHashMap(hashMap, "f", "f1");
    InsertHashMap(hashMap, "g", "g1");
    InsertHashMap(hashMap, "h", "h1");
    PrintHashMap(hashMap);

    int code = RemoveHashMap(hashMap, "a");
    if (code) {
        printf("删除成功\n");
    }
    
    PrintHashMap(hashMap);

    char* res = GetHashMap(hashMap, "g");
    if (res) {
        printf("找到, g -> %s\n", res);
    }
    res = GetHashMap(hashMap, "q");
    if (res == NULL) {
        printf("未找到, q -> %s\n", res);
    }
    
    DeleteHashMap(hashMap);
    printf("销毁完成\n");
}

int main(int argc, char const *argv[])
{
    hashMapTest();
    return 0;
}