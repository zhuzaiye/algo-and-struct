// Copyright (c) 2024/7/16 by hzzhu92
//

#include <vector>
#include <algorithm>

using namespace std;

class Solution {
public:
    int first_miss_positive_num(vector<int> &nums) {
        auto n = nums.size();
        for (auto i = 0; i < n; ++i) {
            // 均摊复杂度分析
            while (nums[i] > 0 && nums[i] <= n && nums[nums[i] - 1] != nums[i]) {
                swap(nums[nums[i] - 1], nums[i]);
            }
        }

        for (int i = 0; i < n; ++i) {
            if (nums[i] != i + 1) {
                return i + 1;
            }
        }
        return n + 1;
    }
};
