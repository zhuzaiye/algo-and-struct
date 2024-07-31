// Copyright (c) 2024/7/16 by hzzhu92
// 如果需要保持原有数组元素的顺序? 如何实现

#include <vector>
#include <algorithm>

using namespace std;

class Solution {
public:
    vector<vector<int>> three_sum(vector<int> &nums) {
        auto n = nums.size();
        // 对原有数组进行排序
        sort(nums.begin(), nums.end());
        // 初始化结果
        vector<vector<int>> ans;

        // 枚举a
        for (int first = 0; first < n; ++first) {
            // 如果当前数字大于0, 后续的两数也一定大于0, 结束后续循环
            if (nums[first] > 0) {
                break;
            }
            // 如何和左边的值相同, 说明重复, 跳过
            if (first > 0 && nums[first] == nums[first - 1]) {
                continue;
            }
            // c 对应的指针初始指向数组最右端
            auto third = n - 1;
            int target = -nums[first];
            // 枚举 b
            for (int second = first + 1; second < n; ++second) {
                // 需要和上一次枚举的数不相同
                if (second > first + 1 && nums[second] == nums[second - 1]) {
                    continue;
                }
                // 需要保证b的指针在c的指针的左侧
                while (second < third && nums[second] + nums[third] > target) {
                    --third;
                }
                // 如果指针重合, 随着 b 后续的增加
                // 就不会满足 a+b+c=0 并且b<c, 退出循环
                if (second == third) {
                    break;
                }
                if (nums[second] + nums[third] == target) {
                    ans.push_back({nums[first], nums[second], nums[third]});
                }
            }
        }
        return ans;
    }
};
