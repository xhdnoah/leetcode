package main

// Given nums = [3,2,2,3], val = 3, return length = 2
// with the first two elements of nums being 2.
// 数组元素在内存地址中连续，不能单独删除某个元素，只能覆盖
// 快慢双指针: 快指针 i 遍历时遇到非 val 元素慢指针 j 才自增
// 这样 i 已经跑到 val 元素后面时 j 指向当前首个 val, 两者交换值
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 0
	for i, num := range nums {
		if num != val {
			if i != j && nums[j] == val {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}
	return j
}
