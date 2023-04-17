package main

import . "leetcode/utils"

// Input: nums = [-2,1,-3,4,-1,2,1,-5,4] Output: 6
// Explanation: [4,-1,2,1] has the largest sum = 6.
// DP: dp[i] 表示 [0,i] 区间内各个子区间和的最大值
func maxSubArray_dp(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	n := len(nums)
	dp, res := make([]int, n), nums[0]
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		res = Max(res, dp[i])
	}
	return res
}

// f(i) 代表以第 i 个数结尾的「连续子数组的最大和」
// 在遍历过程中 考虑 nums[i] 1、独立 2、加入 f(i-1) 取决于 f(i-1)+nums[i] 的值
// 得到 f(i)=max{f(i-1)+nums[i], nums[i]} 可只用一个变量维护 f(i-1) 降低空间复杂度
func maxSubArray_fi(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 { // nums[i-1] 存储 f(i-1) 值
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
