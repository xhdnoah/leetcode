package main

import (
	"sort"
)

// Given array nums = [-1, 0, 1, 2, -1, -4],
// A solution set is: [[-1, 0, 1],[-1, -1, 2]]
// 双指针(解决重复问题) + 排序 or 用 map 提前计算好任意 2 个数字之和
// index 从 1 开始，如果遇到相同元素则 start 直接跳至 index - 1
// start end 指针向中间移动，如果发现移动前后元素相同则跳过 continue
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	length, results := len(nums), make([][]int, 0)
	for index, num := range nums {
		if index == 0 {
			continue
		}
		start, end := 0, length-1
		if index > 1 && num == nums[index-1] {
			start = index - 1
		}
		for start < index && end > index {
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			addNum := nums[start] + num + nums[end]
			if addNum == 0 {
				results = append(results, []int{nums[start], num, nums[end]})
				start++
				end--
			} else if addNum > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return results
}

func threeSumTarget(nums []int, target int) [][]int {
	sort.Ints(nums)
	ln := len(nums)
	var res [][]int
	for i, num := range nums {
		tuples := twoSumTarget(nums, i+1, target-num)
		for _, tuple := range tuples {
			tuple = append(tuple, num)
			res = append(res, tuple)
		}
		for i < ln-1 && num == nums[i+1] {
			i++ // 跳过重复元素
		}
	}
	return res
}

func twoSumTarget(nums []int, start int, target int) [][]int {
	lo, hi := start, len(nums)-1
	var res [][]int
	for lo < hi {
		sum := nums[lo] + nums[hi]
		left, right := nums[lo], nums[hi]
		if sum < target {
			for lo < hi && nums[lo] == left {
				lo++
			}
		} else if sum > target {
			for lo < hi && nums[hi] == right {
				hi--
			}
		} else {
			res = append(res, []int{left, right})
			for lo < hi && nums[lo] == left {
				lo++
			}
			for lo < hi && nums[hi] == right {
				hi--
			}
		}
	}
	return res
}
