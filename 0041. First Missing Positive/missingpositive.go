package main

// 给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数
// Input: [3,4,-1,1], Output: 2; Input: nums = [7,8,9,11,12], Output: 1
// 本题的难点在：只能使用常数级别的额外空间 复杂度：O(1)
// 原地哈希 将数组视为哈希表 哈希函数为：f(nums[i]) = nums[i] - 1
// 3 应该放在索引为 2 的地方 4 应该放在索引为 3 的地方
func firstMissingPositive(nums []int) int {
	size := len(nums)
	for i := 0; i < size; i++ {
		// 先判断这个数字是不是索引，然后判断这个数字是否放在了正确的地方
		// nums[i] != nums[nums[i]-1] 说明没有放在正确位置上需要交换
		for 1 <= nums[i] && nums[i] <= size && nums[i] != nums[nums[i]-1] {
			swap(&nums, i, nums[i]-1)
		}
	}
	for i := 0; i < size; i++ {
		if i+1 != nums[i] {
			return i + 1
		}
	}
	return size + 1
}

func swap(nums *[]int, i, j int) {
	(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
}
