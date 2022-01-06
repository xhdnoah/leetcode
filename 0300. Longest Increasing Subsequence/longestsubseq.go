package main

// Input: nums = [10,9,2,5,3,7,101,18] Output: 4
// The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
// 动态规划 dp[i] 表示以 nums[i] 为结尾的最长子序列长度
func lengthOfLIS(nums []int) int {
	ln, res := len(nums), 0
	dp := make([]int, ln)
	for i := range dp {
		dp[i] = 1
	}
	for i := 0; i < ln; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] { // 构成一个上升对
				dp[i] = max(dp[j]+1, dp[i]) // [..., nums[j], nums[i]] 更新 dp[i]
			}
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
