package main

func bubbleSort(arr []int, size int) []int {
	if size == 1 {
		return arr
	}

	for i := 0; i < size-1; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}

	bubbleSort(arr, size-1)
	return arr
}
