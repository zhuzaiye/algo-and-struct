package twoSum.com;

import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

public class Solution {
    //一遍哈希表
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            int complete = target - nums[i];
            if (map.containsKey(complete)) {
                return new int[]{
                        map.get(complete), i
                };
            }
            map.put(nums[i], i);
        }
        throw new IllegalArgumentException("No two sum solution");
    };

    //两次遍历
    public int[] twoSum1(int[] nums, int target) {
        Map<Integer, Integer> map = new HashMap<>();
        //第一遍历，存入map
        for (int i = 0; i < nums.length; i++) {
            map.put(nums[i], i);
        }
        //第二遍检查有没有
        for (int i = 0; i < nums.length; i++) {
            int complete = target - nums[i];
            if (map.containsKey(complete) && map.get(complete) != i) {
                return new int[]{i, map.get(complete)};
            }
        }
        throw new IllegalArgumentException("No two sum solution");
    }

    public static void main(String[] args) {
        int[] arr = new int[]{2, 2, 7, 11, 5};
        int target = 4;
        Solution solution = new Solution();
        System.out.println(Arrays.toString(solution.twoSum(arr, target)));
        System.out.println(Arrays.toString(solution.twoSum1(arr, target)));
    }
}
