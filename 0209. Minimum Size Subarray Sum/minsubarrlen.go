package main

import (
	"math"
)

// Input: target = 7, nums = [2,3,1,2,4,3] Output: 2
// The subarray [4,3] has the minimal length under the problem constraint.
// 滑动窗口（根据当前子序列情况不断调节子序列起始和终止位置以得出期望结果）
// 窗口的结束位置即为遍历数组的指针 关键在于窗口起始位置如何移动
func minSubArrayLen(target int, nums []int) int {
	i, subLen, sum := 0, math.MaxInt32, 0
	for j, num := range nums {
		sum += num
		for sum >= target {
			subLen = min(subLen, j-i+1)
			sum -= nums[i]
			i++ // 精髓之处
		}
	}
	return subLen
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
