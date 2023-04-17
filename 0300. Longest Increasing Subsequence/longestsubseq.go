package main

import . "leetcode/utils"

// Input: nums = [10,9,2,5,3,7,101,18] Output: 4
// The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
// 动态规划 dp[i] 表示以 nums[i] 为结尾的最长子序列长度
func lengthOfLIS_dp(nums []int) int {
	n, res := len(nums), 0
	dp := make([]int, n)
	for i, num := range nums {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < num { // nums[i] 可以接入以 nums[j] 结尾的上升序列
				dp[i] = Max(dp[j]+1, dp[i]) // [..., nums[j], nums[i]] 更新 dp[i]
			}
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

// 动态规划 + 二分查找 维护 tails 列表 tails[k] 为长度 k+1 子序列的队尾元素
// 本质就是不断把较小数往 tails 里塞，虽然不一定是真正的上升序列但长度是对的
// 考虑序列 78912345 遍历到第三位时 tails:789 再遍历到 1 就得把 1 放到合适位置
// 于是在 tails 二分查找「1」变成了 189（若序列此时结束 因为 res 不变依旧输出 3）再遍历到 2 成为 129 然后是 123 直到 12345
func lengthOfLIS_dichotomy(nums []int) int {
	ans, tails := 0, make([]int, len(nums))
	for _, num := range nums {
		lo, hi := search(tails, ans, num)
		tails[lo] = num
		if ans == hi { // 说明 num 比 tails 的元素都大
			ans++
		}
	}
	return ans
}

func search(nums []int, hi, target int) (int, int) {
	lo := 0
	for lo < hi {
		mid := (lo + hi) / 2
		if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo, hi
}
