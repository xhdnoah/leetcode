package main

// Input: nums = [3,2,1] Output: [1,2,3]
// Input: nums = [1,1,5] Output: [1,5,1]
// 找到下一个排列：将较小数与较大数做原地交换
// 较小数下标尽量靠右，较大数尽可能小，交换后将较大数右边的数按升序重新排列
// 1. 从后向前查找第一个相邻升序的元素对 (i,j) 得到下降区间 [j, end)
// 2. 在下降区间从右往左找第一个满足 nums[k] > nums[i] 则 nums[i] nums[k] 分别为较小和较大数
// 3. 交换 nums[i] nums[k], 最后原地交换 [j, end) 区间元素使其从降序变升序
// 特：如果第一步找不到符合条件 i, 说明已经是最大排列，直接执行 3 生成最小排列
func nextPermutation(nums []int) {
	ln := len(nums)
	i, j, k := ln-2, ln-1, ln-1
	for i >= 0 && nums[i] >= nums[j] { // find first nums[i] < nums[j]
		i--
		j--
	}

	if i >= 0 { // 不是最后一个排列
		for nums[i] >= nums[k] { // find nums[i] < nums[k]
			k--
		}
		nums[i], nums[k] = nums[k], nums[i]
	}
	// reverse nums[j:end]
	for i, j := j, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// 对升/降序数组原地反转
func reverse(nums *[]int, i, j int) {
	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}

func swap(nums *[]int, i, j int) {
	(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
}
