#include <stdlib.h>

// 快速排序
void QuickSort(int* a, int low, int high) {
    if (low > high) {
        return;
    }
}

// 比较大小, 便于qsort的从小到大排序
int comp(void* a, void* b) { return *(int*)a > *(int*)b ? 1 : -1; }

// 双指针法
int** threeSum(int* nums, int numsSize, int* returnSize, int** returnColumnSizes) {
    qsort(nums, numsSize, sizeof(int), comp);
    // 第一次过滤
    if (numsSize < 3 || nums[0] > 0 || nums[0] + nums[1] + nums[2] > 0) {
        return NULL;
    }

    int** rst = NULL;
    int size = 0, left = 0, right = 0;
    for (int i = 0; i < numsSize - 2 && nums[i] <= 0; i++) {
        left = i + 1;
        right = numsSize - 1;
        while (left < right) {
            // 双指针移动
            if (nums[i] + nums[left] + nums[right] > 0) {
                right--;
            } else if (nums[i] + nums[left] + nums[right] < 0) {
                left++;
            } else {
                size++;
                rst = (int**)realloc(rst, sizeof(int*) * size);
                rst[size - 1] = (int*)malloc(sizeof(int) * 3);
                rst[size - 1][0] = nums[i];
                rst[size - 1][1] = nums[left];
                rst[size - 1][2] = nums[right];
                // 去重
                while (left + 1 < numsSize && nums[left] == nums[left + 1])
                    left++;
                while (right - 1 >= 0 && nums[right] == nums[right - 1])
                    right--;
                left++;
                right--;
            }
        }
        // 去重
        while (i + 1 < numsSize && nums[i] == nums[i + 1]) {
            i++;
        }
    }
    *returnSize = size;
    *returnColumnSizes = malloc(sizeof(int) * size);
    return rst;
}

/*
    qsort：
        base -- 指向要排序的数组的第一个元素的指针。
        nitems -- 由 base 指向的数组中元素的个数。
        size -- 数组中每个元素的大小，以字节为单位。
        compar -- 用来比较两个元素的函数。
 */