package main

// 归并排序 分治+后序遍历
func mergeSort(arr []int) []int {
	// 递归终止条件
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	// 先递归使两个子序列有序 再对其合并
	left := mergeSort(arr[0:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

// 假定 left right 已经有序 使用双指针进行线性合并
func merge(left, right []int) (res []int) {
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}
	// 处理 left right 中剩余元素
	res = append(res, left[i:]...)
	res = append(res, right[j:]...)
	return
}
