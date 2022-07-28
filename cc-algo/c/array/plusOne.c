/*
   加一：
       给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
       最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
       你可以假设除了整数 0 之外，这个整数不会以零开头。
*/
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* plusOne2(int* digits, int digitsSize, int* returnSize) {
  for (int i = digitsSize - 1; i >= 0; --i) {
    if (digits[i] == 9) {
      digits[i] = 0;
    } else {
      digits[i]++;
      *returnSize = digitsSize;
      return digits;
    }
  }
  int* result = (int*)malloc(sizeof(int) * (digitsSize + 1));
  memset(result, 0, (digitsSize + 1) * sizeof(int));
  result[0] = 1;
  *returnSize = digitsSize + 1;
  return result;
}

int* plusOne(int* digits, int digitsSize, int* returnSize) {
  // malloc the memerisize
  // malloc仅仅是分配一块连续的内存，且未初始化
  // free
  // calloc
  // realloc
  // void *类型可以强转为任何其他类型的的指针

  for (int i = digitsSize - 1; i >= 0; i--) {
    if (digits[i] < 9) {
      digits[i] = digits[i] + 1;
      *returnSize = digitsSize;
      return digits;
    } else {
      digits[i] = 0;
    }
  }
  // 如果
  int* mer = (int*)malloc((digitsSize + 1) * sizeof(int));
  mer[0] = 1;
  *returnSize = digitsSize + 1;
  return mer;
}

void arryprint(int* arry, int* size) {
  for (int i = 0; i < (*size); i++) {
    printf("%d ", arry[i]);
  }
  printf("\n");
}

int main() {
  /* code */
  int arr1[] = {1, 2, 4};
  int arr2[] = {1, 2, 9};
  int arr3[] = {9};
  int* size1 = (int*)malloc(sizeof(int));
  int* rn1 = plusOne(arr1, 3, size1);
  arryprint(rn1, size1);

  int* size2 = (int*)malloc(sizeof(int));
  int* rn2 = plusOne(arr2, 3, size2);
  arryprint(rn2, size2);

  int* size3 = (int*)malloc(sizeof(int));
  int* rn3 = plusOne(arr3, 1, size3);
  arryprint(rn3, size3);

  system("pause");
  return 0;
}
