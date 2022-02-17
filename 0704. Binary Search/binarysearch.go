package main

// Input: nums = [-1,0,3,5,9,12], target = 9 Output: 4
// 区间的定义就是不变量，即要在 二分查找 的过程中保持不变量
// 就是在 while 中每次边界处理都坚持根据区间定义操作，这就是循环不变量规则
// 定义 target 在一个左闭右闭区间 [low, high]
func search_closed(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high { // 闭区间下 low == high 有存在意义
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1 // 更新为 mid-1 因当前 nums[mid] 一定不是 target 而下个查找区间也是闭区间
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 左闭右开区间 [low, high)
func search_open(nums []int, target int) int {
	low, high := 0, len(nums)
	for low < high { // 用 < 因为 low==high 在开区间下没有意义
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid // 下个寻找区间也是左闭右开区间，故更新为 mid 即不会去比较 nums[mid]
		} else {
			low = mid + 1
		}
	}
	return -1
}
