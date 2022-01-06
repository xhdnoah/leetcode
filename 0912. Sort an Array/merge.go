package main

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[0:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	leftLen, rightLen := len(left), len(right)
	res := make([]int, leftLen+rightLen)

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
	for ; i < leftLen; i++ {
		res = append(res, left[i])
	}
	for ; j < rightLen; j++ {
		res = append(res, right[j])
	}
	return res
}
