// 复杂度 O(n)
int searchInsert(int* nums, int numsSize, int target) {
    for (int i = 0; i < numsSize; i++) {
        if (target <= nums[i]) {
            return i;
        }
    }
    return numsSize;
}

int searchInsertBinary(int* nums, int numsSize, int target) {
    if (nums[numsSize - 1] <= target) {
        return numsSize;
    }
    int left = 0, right = numsSize - 1;
    while (left < right) {
        int mid = left + (left + right) / 2;
        if (nums[mid] < target) {
            left = mid + 1;
        } else {
            right = mid;
        }
    }
    return left;
}