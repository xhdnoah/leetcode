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
	switch pivot := randomPartition(nums, low, high); {
	case pivot < target:
		return quickSelect(nums, pivot+1, high, target)
	case pivot > target:
		return quickSelect(nums, low, pivot-1, target)
	default:
		return nums[pivot]
	}
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
	// 在之前遍历过程中，满足 nums[low+1,..i] < pivot 且 nums[i+1,..j] >= pivot
	nums[low], nums[i] = nums[i], nums[low]
	// 交换后 nums[low..i-1] < pivot, nums[i] = pivot, nums[i+1..high] >= pivot
	return i
}

// 基于堆排序的选择方法
// 建立一个大根堆 做 k−1 次删除操作后堆顶元素就是答案
func findKthLargest_heap(nums []int, k int) int {
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}
	return nums[0]
}

func buildMaxHeap(nums []int, heapSize int) {
	for i := heapSize / 2; i >= 0; i-- {
		maxHeapify(nums, i, heapSize)
	}
}

func maxHeapify(nums []int, i, heapSize int) {
	l, r, largest := i*2+1, i*2+2, i
	if l < heapSize && nums[l] > nums[largest] {
		largest = l
	}
	if r < heapSize && nums[r] > nums[largest] {
		largest = r
	}
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		maxHeapify(nums, largest, heapSize)
	}
}
