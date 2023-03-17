package main

// Input: [2, 3, 1, 0, 2, 5, 3] Output: 2 or 3
// 数组元素的索引和值是一对多的关系 交换操作使索引和值一一对应 nums[i]==i
// 遍历中，第一次遇到 x 时将其交换至索引 x 处 而当第二次遇到 x 一定有 nums[x]==x 即可得到一组重复数字

func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == i {
			continue
		}
		if nums[nums[i]] == nums[i] {
			return nums[i]
		}
		nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
	}
	return -1
}
