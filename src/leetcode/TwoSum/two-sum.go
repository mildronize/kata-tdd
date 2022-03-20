/*
Problem: https://leetcode.com/problems/two-sum
Author: Thada Wangthammang
*/

package two_sum

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{0, 1}
}
