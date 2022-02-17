package main

// Input: nums = [5,7,7,8,8,10], target = 8 Output: [3,4]
// 二分查找：由于数组已排序是单调递增的，可以利用二分法来加速查找过程
// 寻找 leftIdx 即为在数组中寻找第一个大于等于 target 的下标
// 寻找 rightIdx 即为在数组中寻找第一个大于 target 的下标，然后将下标减一
// 两者判断条件不同，定义 binarySearch(nums, target, lower) 表示在 nums 数组中二分查找 target 的位置
// 如果 lower 为 true，则查找第一个大于等于 target 的下标，否则查找第一个大于 target 的下标
func binarySearch(nums []int, target int, lower bool) int {
	left, right, res := 0, len(nums)-1, len(nums)
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target || (lower && nums[mid] >= target) {
			right = mid - 1
			res = mid // 不断更新 res 直至找到最左边那个大于（等于）target 的下标
		} else {
			left = mid + 1
		}
	}
	return res
}

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	leftIdx := binarySearch(nums, target, true)
	rightIdx := binarySearch(nums, target, false) - 1
	if leftIdx <= rightIdx && rightIdx < len(nums) && nums[leftIdx] == target && nums[rightIdx] == target {
		res = []int{leftIdx, rightIdx}
	}
	return res
}
