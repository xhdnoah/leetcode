package main

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	cur, res := make([]int, 0), make([][]int, 0)
	sort.Ints(nums)
	kSum(nums, 4, 0, target, cur, &res)
	return res
}

// nums 已排过序
func kSum(nums []int, k, left, target int, cur []int, res *[][]int) {
	ln := len(nums)
	// 边界条件过滤
	if ln-left < k || k < 2 || nums[left]*k > target || nums[ln-1]*k < target {
		return
	}
	if k == 2 {
		twoSum(nums, left, target, cur, res)
	} else {
		for i := left; i < ln; i++ {
			// 跳过重复项
			if i == left || (i > left && nums[i] != nums[i-1]) {
				next := make([]int, len(cur))
				// cur 内容不能改变，还会在其他递归线中用到
				copy(next, cur)
				next = append(next, nums[i])
				kSum(nums, k-1, i+1, target-nums[i], next, res)
			}
		}
	}
}

func twoSum(nums []int, left, target int, cur []int, res *[][]int) {
	right := len(nums) - 1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			tmp = append(tmp, nums[left], nums[right])
			*res = append(*res, tmp)
			left++
			right--
			// 跳过重复项
			for left < right && nums[left] == nums[left-1] {
				left++
			}
			for left < right && nums[right] == nums[right+1] {
				right--
			}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
}
