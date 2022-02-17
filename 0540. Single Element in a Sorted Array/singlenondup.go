package main

// Input: nums = [1,1,2,3,3,4,4,8,8] Output: 2
func singleNonDuplicate(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		// 二分的核心是确定左右指针变化的条件
		// 在单一元素的左侧 相邻的奇偶元素相等
		// i 为偶数 nums[i] == nums[i+1]; i 为奇数 nums[i] == nums[i-1]
		// 取 mid 如满足上述则将区间左侧收缩到 mid+1 否则收缩右侧到 mid
		if nums[mid] == nums[mid^1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}
