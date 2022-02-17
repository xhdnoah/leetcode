package main

import (
	"math/rand"
)

// 快排 partition 得到 p 点的下标即为最终排序后的下标，根据下标可以判断第 K 大的数位置
// 时间复杂度 O(n)，空间复杂度 O(log n)，最坏时间复杂度为 O(n^2)，空间复杂度 O(n)
func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, low, high, target int) int {
	pivot := randomPartition(nums, low, high)
	if pivot == target {
		return nums[pivot]
	} else if pivot < target {
		return quickSelect(nums, pivot+1, high, target)
	}
	return quickSelect(nums, low, pivot-1, target)
}

func randomPartition(nums []int, low, high int) int {
	// 随机初始化 pivot 元素避免极端用例
	i := low + rand.Intn(high-low+1)
	nums[i], nums[low] = nums[low], nums[i]
	return partition(nums, low, high)
}

func partition(nums []int, low, high int) int {
	i, pivot := low, nums[low] // 游标 i 确定最终基准点
	for j := low + 1; j <= high; j++ {
		if nums[j] < pivot { // j 向后查找小于 pivot 数
			i++ // i 初值为 low 先右移再交换，小于 pivot 的元素都被交换到前面
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	// 在之前遍历过程中，满足 nums[low+1..i] < pivot 且 nums[i..j] >= pivot
	nums[low], nums[i] = nums[i], nums[low]
	// 交换后 nums[low..i-1] < pivot, nums[i] = pivot, nums[i+1..high] >= pivot
	return i
}
