package main

import (
	"sort"
)

// Given array nums = [-1, 0, 1, 2, -1, -4],
// A solution set is: [[-1, 0, 1],[-1, -1, 2]]
// 难点在于去重，双指针(解决重复问题) + 排序 or 用 map 提前计算好任意 2 个数字之和
// 固定住最左最小数字指针 index 双指针 start end 分设在数组索引 [index+1:len(nums)-1] 两端
// 通过双指针交替向中间移动，记录对于每个固定指针 index 的所有满足三数和为 0 的 start end 组合
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n, ans := len(nums), make([][]int, 0)

	for index, num := range nums[:n-1] {
		if index > 0 && num == nums[index-1] {
			continue // 去重 skip the same nums[index] 此时已经将 nums[index-1] 所有组合加入到结果
		}
		start, end := index+1, n-1
		for start < end { // double pointer
			switch sum := num + nums[start] + nums[end]; {
			case sum < 0:
				start++ // 左指针右移同时去重
				for start < end && nums[start] == nums[start-1] {
					start++
				}
			case sum > 0:
				end-- // 右指针左移同时去重
				for start < end && nums[end] == nums[end+1] {
					end--
				}
			default:
				ans = append(ans, []int{num, nums[start], nums[end]})
				start++
				end--
				for start < end && nums[start] == nums[start-1] {
					start++
				}
				for start < end && nums[end] == nums[end+1] {
					end--
				}
			}
		}
	}
	return ans
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
