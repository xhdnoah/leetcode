package main

import (
	"container/heap"
	"sort"
)

// Input: nums = [1,3,-1,-3,5,3,6,7], k = 3 Output: [3,3,5,5,6,7]
// Explanation:
// Window position                Max
// ---------------               -----
// [1  3  -1] -3  5  3  6  7       3
//  1 [3  -1  -3] 5  3  6  7       3
//  1  3 [-1  -3  5] 3  6  7       5
//  1  3  -1 [-3  5  3] 6  7       5
//  1  3  -1  -3 [5  3  6] 7       6
//  1  3  -1  -3  5 [3  6  7]      7
// 对于「最大值」有非常合适的数据结构是优先队列（堆）其中大根堆可以实时维护一系列元素中的最大值
// 优先队列中的每个元素都有各自的优先级，优先级最高的元素最先得到服务(出队，优先队列往往用堆来实现。
// 堆是一种特别的完全二叉树 若母节点的值恒小于等于子节点的值，则称为最小堆 min heap 反之为最大堆 max heap
// 堆化就是顺着节点所在的路径，向上或者向下，对比交换 要实现一个 heap，需要实现 5个抽象方法 Len Less Swap Push Pop
// 对于本题，初始时将数组前 k 个元素放入优先队列。每当窗口向右移动将一个新元素放入优先队列，此时堆顶元素就是堆中所有元素的最大值
// 然而这个最大值可能并不在滑动窗口中，这种情况下这个值在数组中的位置出现在滑动窗口左边界的左侧，之后不会再碰到可以永远从队列移除
// 不断地移除堆顶的元素，直到其确实出现在滑动窗口中。此时堆顶元素就是滑动窗口中的最大值。为方便判断堆顶元素与滑动窗口的位置关系，可在优先队列中存储二元组 (num, index)
var a []int

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool { return a[h.IntSlice[i]] > a[h.IntSlice[j]] } // 大根堆 or 小根堆
// 在调用自定义 Push Pop 时 heap 源码中已经做了相应的上浮(up)下沉(down)处理确保堆的有序性，使用者只需实现最简单的追加、删除操作
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func maxSlidingWindow(nums []int, k int) []int {
	a = nums
	priorityQueue := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		priorityQueue.IntSlice[i] = i
	}
	heap.Init(priorityQueue)

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[priorityQueue.IntSlice[0]]
	for i := k; i < n; i++ {
		heap.Push(priorityQueue, i)
		for priorityQueue.IntSlice[0] <= i-k {
			heap.Pop(priorityQueue)
		}
		ans = append(ans, nums[priorityQueue.IntSlice[0]])
	}
	return ans
}
