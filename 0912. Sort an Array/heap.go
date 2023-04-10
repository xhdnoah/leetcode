package main

// 堆排序的思想就是先将待排序的序列建成大根堆
// 使得每个父节点元素 >= 其子节点。此时整个序列最大值即为堆顶元素
// 将其与末尾元素交换，使末尾元素为最大值
// 然后再调整堆顶元素使剩下的 n−1 个元素仍为大根堆得到次大值
// 重复执行以上操作得到一个从顶到尾有序的序列
// 因为堆是完全二叉树因此可以用数组存储
func heapSort(arr []int) {
	n := len(arr)
	// 构建初始最大堆
	for i := n/2-1; i >= 0; i-- {
		heapify(arr, n, i)
	}
	// 逐步取出堆顶元素 调整剩余元素为最大堆
	for i := n-1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		// 对剩余元素进行最大堆调整
		heapify(arr, i, 0)
	}
}

// 对 [4,6,8,5,9] 建堆得到 [9,6,8,5,4]
//   9
//  6 8
// 5 4
func heapify(arr []int, n, i int) {
	// 目的找到当前节点、左右子节点中的最大值
	left, right, largest := 2*i+1, 2*i+2, i 
	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i] // 将最大值与当前节点交换
		heapify(arr, n , largest) // 递归地对交换后的节点进行最大堆调整
	}
}