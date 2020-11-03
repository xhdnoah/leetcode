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
