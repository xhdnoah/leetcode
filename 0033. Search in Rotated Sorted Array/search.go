package main

// 整数数组 nums 按升序排列在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行旋转
// 使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）
// 例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 在旋转数组中搜索一个给定的目标值
// Input: nums = [4,5,6,7,0,1,2], target = 0 Output: 4 算法时间复杂度必须是 O(log n)
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target { // 说明
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right + 1
}
