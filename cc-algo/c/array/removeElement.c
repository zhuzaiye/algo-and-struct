#include <stdio.h>
#include <stdlib.h>

// 双指针或者快慢指针，记录位置和推动数据移动
// nums 整型数组
// numSize nums数组的参长度
// val 需要移除的对象整数值
int removeElement(int* nums, int numsSize, int val) {
    int lp = 0;
    for (int i = 0; i < numsSize; i++) {
        if (nums[i] != val) {
            nums[lp] = nums[i];
            lp++;
        }
    }
    return lp;
}

int removeElement2(int* nums, int numsSize, int val) {
    int fast = 0, slow = 0;
    while (fast < numsSize) {
        if (nums[fast++] != val) {
            nums[slow++] = nums[fast - 1];
        }
    }
    return slow;
}

void foreachArray(int* pArr, int len) {
    int i;
    for (i = 0; i < len; i++) {
        printf(" %d", pArr[i]);
    }
    printf(" \n");
}

int main(int argc, char const* argv[]) {
    int nums[] = {3, 2, 2, 3};
    int len = removeElement(nums, 4, 2);
    printf("长度： %d\n", len);
    foreachArray(nums, 2);
    return 0;
}
