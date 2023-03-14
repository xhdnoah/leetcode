package main

import . "leetcode/utils"

// Input: nums = [-2,1,-3,4,-1,2,1,-5,4] Output: 6
// Explanation: [4,-1,2,1] has the largest sum = 6.
// DP: dp[i] 表示 [0,i] 区间内各个子区间和的最大值
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	ln := len(nums)
	dp, res := make([]int, ln), nums[0]
	dp[0] = nums[0]
	for i := 1; i < ln; i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		res = Max(res, dp[i])
	}
	return res
}
