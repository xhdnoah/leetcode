package main

// Given nums = [3,2,2,3], val = 3, return length = 2
// with the first two elements of nums being 2.
// 数组元素在内存地址中连续，不能单独删除某个元素，只能覆盖
// 快慢双指针: i 遍历查找到非 val 元素与指向 val 的 j 交换值并自增 j 直至找到下一个 val
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
