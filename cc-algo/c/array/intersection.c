#include <stdio.h>
#include <stdlib.h>

#include "../util/src/uthash.h"

struct hashTable {
    int key; /* key int型 */
    int val; /* value int型 */
    UT_hash_handle hh;
};

int* inetersection(int* nums1, int nums1Size, int* nums2, int nums2Size,
                   int* returnSize) {
    if ((nums1 == NULL) || (nums2 == NULL) || (nums2Size <= 0) ||
        (nums1Size == 0)) {
        *returnSize = 0;
        return NULL;
    }

    struct hashTable* entry = NULL;
    for (int i = 0; i < nums1Size; i++) {
        struct hashTable* tmp = NULL;
        HASH_FIND_INT(entry, &nums1[i], tmp);
        if (tmp == NULL) {
            struct hashTable* node = malloc(sizeof(struct hashTable));
            node->key = nums1[i], node->val = 0;
            HASH_ADD_INT(entry, key, node);
        }
    }

    int minIntersectionLen = nums1Size <= nums2Size ? nums1Size : nums2Size;
    int* rst = (int*)malloc(sizeof(int) * minIntersectionLen);
    int index = 0;
    for (int i = 0; i < nums2Size; i++) {
        struct hashTable* tmp = NULL;
        HASH_FIND_INT(entry, &nums2[i], tmp);
        if ((tmp != NULL) && (tmp->val == 0)) {
            rst[index++] = tmp->key;
            tmp->val = 1;
        }
    }

    // free malloc ??
    /* 
        segmention-fault
        关于这个错误如何排查？如何解决？
        本次尝试解决的就在line49 -- line55之间
     */
    struct hashTable *temp, *iter;
    for (iter = entry, temp = entry; iter != NULL; iter = temp) {
        temp = temp->hh.next;
        HASH_DEL(entry, iter);  // 先从hash表中删除
        free(iter);                 // 然后释放该节点
    }
    entry = NULL;
    *returnSize = index;
    return rst;
}

void arryprint(int* arry, int* size) {
    for (int i = 0; i < (*size); i++) {
        printf("%d ", arry[i]);
    }
    printf("\n");
}

void main() {
    int arr1[] = {1, 2, 2, 1};
    int arr2[] = {2, 2};
    int* size1 = (int*)malloc(sizeof(int));
    int* rn = inetersection(arr1, 4, arr2, 2, size1);
    arryprint(rn, size1);
}
