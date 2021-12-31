package main

// Input: [2,0,2,1,1,0] Output: [0,0,1,1,2,2]
// A rather straight forward solution is a two-pass algorithm using counting sort.
// First, iterate the array counting number of 0’s, 1’s, and 2’s, then overwrite array with total number of 0’s, then 1’s and followed by 2’s.
// Could you come up with a one-pass algorithm using only constant space?
// 计数排序，适合待排序数字很少，用 3 个容器分别记录 0，1，2 出现个数再根据个数排列 0，1，2 即可，时间复杂度 O(n)，空间复杂度 O(K)
// 游标移动控制顺序，0 排在最前面，所以只要添加一个 0，就需要放置 1 和 2。1 排在 2 前面，所以添加 1 的时候也需要放置 2 。至于最后的 2，只用移动游标即可。
func sortColors(nums []int) {
	zero, one := 0, 0
	for i, n := range nums {
		nums[i] = 2
		if n <= 1 {
			nums[one] = 1
			one++
		}
		if n == 0 {
			nums[zero] = 0
			zero++
		}
	}
}
