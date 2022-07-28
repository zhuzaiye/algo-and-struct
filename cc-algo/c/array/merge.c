#include <stdio.h>

void merge(int* nums1, int nums1Size, int m, int* nums2, int nums2Size, int n) {
    int i = m - 1, j = n - 1, len = nums1Size - 1;
    while (j >= 0 && i >= 0) {
        if (nums2[j] <= nums1[i]) {
            nums1[len--] = nums1[i--];
        } else {
            nums1[len--] = nums2[j--];
        }
    }

    while (j >= 0) {
        nums1[len--] = nums2[j--];
    }
}

void merge2(int* nums1, int nums1Size, int m, int* nums2, int nums2Size,
            int n) {
    int i = m - 1, j = nums2Size - 1, k = nums1Size - 1;
    while (i != k && j >= 0) {
        if (i != -1 && nums1[i] > nums2[j]) {
            nums1[k--] = nums1[i--];
        } else {
            nums1[k--] = nums2[j--];
        }
    }
}

void arryprint(int* arry, int size) {
    for (int i = 0; i < size; i++) {
        printf("%d ", arry[i]);
    }
    printf("\n");
}

int main() {
    /* code */
    int arr1[] = {1, 2, 3, 0, 0, 0};
    int arr2[] = {1, 2, 9};
    merge(arr1, 6, 3, arr2, 3, 3);
    arryprint(arr1, 6);
    system("pause");
    return 0;
}

/* 
 (shift + alt + a) Toggle Comment for VSCode IDE.

 主要想法：
    1. nums1 内存长度已存在，cap=len(nums1)+len(nums2)
    2. nums1 和 num2 已经排好序
    3. 将nums1和nums2从内容实际长度位置，从后往前进行比较添加到nums的cap
    4. 处理没有遍历完的那一方对nums1剩余的空位进行填充，实际就是倒序添加即可
 */