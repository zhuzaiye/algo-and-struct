// Copyright (c) 2024/7/14 by hzzhu92
// Using unordered_map

#include <vector>
#include <unordered_map>

using namespace std;

class TwoSumSolution {
public:
    vector<int> two_sum(vector<int> &nums, int target) {
        unordered_map<int, int> hash_table;
        for (int i = 0; i < nums.size(); i++) {
            auto it = hash_table.find(target - nums[i]);
            if (it != hash_table.end()) {
                return {it->second, i};
            }
            hash_table[nums[i]] = i;
        }
        return {};
    }
};
