#! usr/bin/env python
# _*_ coding: utf-8 _*_

"""
Date: 2020/8/5 7:05
Auth: JoeLang
Desc:
"""

from typing import List


def two_sum(nums: List[int], target: int) -> List[int]:
    couple = {}
    for index, num in enumerate(nums):
        if target - num in couple.keys():
            return [nums.index(target - num), index]
        else:
            couple[num] = target - num
    return [-1, -1]


if __name__ == '__main__':
    arr = [2, 2, 7, 11,  5]
    tar = 4
    print(two_sum(arr, tar))
