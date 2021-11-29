package main

// Input: nums = [3,2,1] Output: [1,2,3]
// Input: nums = [1,1,5] Output: [1,5,1]
// 找到下一个排列：将较小数与较大数做原地交换
// 较小数下标尽量靠右，较大数尽可能小，交换后将较大数右边的数按升序重新排列
// 1. 从右往左遍历找到第一个满足 nums[i] < nums[i+1] 此时较小数为 nums[i] 且 [i+1, n) 为下降区间
// 2. 在下降区间从右往左找第一个满足 nums[i] < nums[j], 此时较大数为 nums[j]
// 3. 交换 nums[i] nums[j], 最后原地交换 [i+1, n) 区间元素使其变为升序
// 特：如果第一步找不到符合条件 i, 说明已经是最大排列，直接执行 3 生成最小排列
func nextPermutation(nums []int) {
	ln := len(nums)
	i, j := ln-2, ln-1
	for ; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i >= 0 {
		for ; j > i; j-- {
			if nums[j] > nums[i] {
				break
			}
		}
		swap(&nums, i, j)
	}
	reverse(&nums, i+1, ln-1)
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
