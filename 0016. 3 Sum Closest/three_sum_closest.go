package main

import (
	"math"
	"sort"
)

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

// Given array nums = [-1, 2, 1, -4], and target = 1.
// The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
// 双指针夹逼，j 为 i 下一个数字，k 为数组最后一个数字，逐渐夹出最近值
// 重复问题 1. map 计数去重 2. i 在循环中与前一位比较，如果相等则后移至不相等
func threeSumClosest(nums []int, target int) int {
	n, res, diff := len(nums), 0, math.MaxInt32
	if n > 2 {
		sort.Ints(nums)
		for i := 0; i < n-2; i++ {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			for j, k := i+1, n-1; j < k; {
				sum := nums[i] + nums[j] + nums[k]
				if abs(sum-target) < diff {
					res, diff = sum, abs(sum-target)
				}
				if sum == target {
					return sum
				}
				if sum > target {
					k--
				} else {
					j++
				}
			}
		}
	}
	return res
}
