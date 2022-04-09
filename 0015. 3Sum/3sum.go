package main

import (
	"sort"
)

// Given array nums = [-1, 0, 1, 2, -1, -4],
// A solution set is: [[-1, 0, 1],[-1, -1, 2]]
// 难点在于去重，双指针(解决重复问题) + 排序 or 用 map 提前计算好任意 2 个数字之和
// 固定住最左（小）数字指针 index，双指针 start end 分设在数组索引 (index, len(nums)) 两端
// 通过双指针交替向中间移动，记录对于每个固定指针 index 的所有满足三数和为 0 的 start end 组合
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ln, res := len(nums), make([][]int, 0)
	if ln == 0 {
		return res
	}
	for index := range nums[:ln-1] {
		if nums[index] > 0 {
			break // because end > start > index 三个数字都大于 0
		}
		if index > 0 && nums[index] == nums[index-1] {
			continue // skip the same nums[index] 此时已经将 nums[index-1] 所有组合加入到结果
		}
		start, end := index+1, ln-1
		for start < end { // double pointer
			sum := nums[index] + nums[start] + nums[end]
			if sum < 0 {
				start += 1
				for start < end && nums[start] == nums[start-1] {
					start += 1
				}
			} else if sum > 0 {
				end -= 1
				for start < end && nums[end] == nums[end+1] {
					end -= 1
				}
			} else {
				res = append(res, []int{nums[index], nums[start], nums[end]})
				start += 1
				end -= 1
				for start < end && nums[start] == nums[start-1] {
					start += 1
				}
				for start < end && nums[end] == nums[end+1] {
					end -= 1
				}
			}
		}
	}
	return res
}

func threeSumTarget(nums []int, target int) [][]int {
	sort.Ints(nums)
	ln := len(nums)
	var res [][]int
	for i := 0; i < ln; i++ {
		tuples := twoSumTarget(nums, i+1, target-nums[i])
		for _, tuple := range tuples {
			tuple = append(tuple, nums[i])
			res = append(res, tuple)
		}
		for i < ln-1 && nums[i] == nums[i+1] {
			i++ // 跳过重复元素
		}
	}
	return res
}

// 双指针
func twoSumTarget(nums []int, start int, target int) [][]int {
	lo, hi := start, len(nums)-1
	var res [][]int
	for lo < hi {
		sum := nums[lo] + nums[hi]
		left, right := nums[lo], nums[hi]
		if sum < target {
			for lo < hi && nums[lo] == left { // 跳过重复元素
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
