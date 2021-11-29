package main

// Given nums = [3,2,2,3], val = 3, return length = 2
// with the first two elements of nums being 2.
// 双指针 i 遍历查找下一个非 val 元素交换到 j 处并自增 j, j 会找到第一个 val 停住
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
