package main

// Input: nums = [1,3,5,6], target = 5 Output: 2; target = 2 Output: 1
// 在排序数组中寻找一个目标值，可以联想到二分法在 O(logn) 时间内找到目标值是否存在
// 考虑插入位置 pos 成立条件为 nums[pos-1] < target <= nums[pos] 即找到第一个大于等于 target 的下标（二分逼近
func searchInsert(nums []int, target int) int {
	lo, hi, res := 0, len(nums)-1, len(nums)
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] >= target {
			hi = mid - 1
			res = mid
		} else {
			lo = mid + 1
		}
	}
	return res
}
