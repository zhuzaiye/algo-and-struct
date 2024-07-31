// Copyright (c) 2024/7/14 by hzzhu92
//

#define DOCTEST_CONFIG_IMPLEMENT
#include "doctest.h"

//#include <vector>
//#include <ostream>

#include "hash_map/two_sum.cpp"

//namespace std {
//ostream &operator<<(ostream &os, const vector<int> &vec) {
//    os << "[";
//    for (size_t i = 0; i < vec.size(); ++i) {
//        if (i != 0) {
//            os << ", ";
//        }
//        os << vec[i];
//    }
//    os << "]";
//    return os;
//}
//}

TEST_CASE("Test two_sum method") {
    TwoSumSolution sol;

    SUBCASE("Test case 1") {
        vector<int> nums = {2, 7, 11, 15};
        int target = 9;
        vector<int> result = sol.two_sum(nums, target);
        vector<int> expected = {0, 1};

        CHECK(result == expected);
    }
}