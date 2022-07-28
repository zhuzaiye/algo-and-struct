#include <stdio.h>
#include <stdlib.h>

#include "../util/src/uthash.h"

int* twoSum(int* nums, int numsSize, int target, int* returnSize) {
    int i, j;
    int* rst = NULL;
    *returnSize = 0;
    for (i = 0; i < numsSize - 1; i++) {
        for (j = i + 1; j < numsSize; j++) {
            if (nums[i] + nums[j] == target) {
                rst = (int*)malloc(sizeof(int) * 2);
                rst[0] = i;
                rst[1] = j;
                *returnSize = 2;
                return rst;
            }
        }
    }
    return rst;
}

void arryprint(int* arry, int* size) {
    for (int i = 0; i < (*size); i++) {
        printf("%d ", arry[i]);
    }
    printf("\n");
}

// UT Hash 的使用
struct hashTable {
    int key; /* key int型 */
    int val; /* value int型 */
    UT_hash_handle hh;
};

struct hashTable* hashtable;

struct hashTable* find(int ikey) {
    struct hashTable* tmp;
    HASH_FIND_INT(hashtable, &ikey, tmp);
    return tmp;
}

void insert(int ikey, int ival) {
    struct hashTable* it = find(ikey);
    if (it == NULL) {
        struct hashTable* tmp = malloc(sizeof(struct hashTable));
        tmp->key = ikey, tmp->val = ival;
        HASH_ADD_INT(hashtable, key, tmp);
    } else {
        it->val = ival;
    }
}

int* twoSum2(int* nums, int numsSize, int target, int* returnSize) {
    int* rst = NULL;
    if ((nums == NULL) || (numsSize == 0)) {
        *returnSize = 0;
        return rst;
    }

    int sub;
    for (int i = 0; i < numsSize; i++) {
        sub = target - nums[i];
        struct hashTable* node = find(sub);
        if (node == NULL) {
            insert(nums[i], i);
        } else {
            rst = (int*)malloc(sizeof(int) * 2);
            rst[0] = node->val;
            rst[1] = i;
            *returnSize = 2;
            return rst;
        }
    }
    return rst;
}

void main() {
    int nums[] = {2, 11, 7, 15};
    int target = 9;
    int* size1 = (int*)malloc(sizeof(int));
    int* rn = twoSum(nums, 4, target, size1);
    arryprint(rn, size1);

    printf("------------------\n");
    int* rn2 = twoSum2(nums, 4, target, size1);
    arryprint(rn2, size1);
}