package main

func partition(arr []int, low, high int) int {
	pivot := arr[low]
	for low < high {
		for low < high && arr[high] >= pivot {
			high-- // 从右往左找第一个小于基准的数赋值给 arr[low]
		}
		arr[low] = arr[high]
		for low < high && arr[low] <= pivot {
			low++ // 从左往右找第一个大于基准的数赋值给 arr[high]
		}
		arr[high] = arr[low]
	}
	arr[low] = pivot
	return low // 最终基准点
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		pivot := partition(arr, low, high)
		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}
	return arr
}
