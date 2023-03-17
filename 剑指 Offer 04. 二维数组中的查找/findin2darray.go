package main

// 由升序排列 可以对每一行使用一次二分查找
func findNumberIn2DArray(martix [][]int, target int) bool {
	for _, row := range martix {
		i := binarySearch(row, target)
		if row[i] == target {
			return true
		}
	}
	return false
}

func binarySearch(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		val := nums[mid]
		if val == target {
			return mid
		} else if val > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return -1
}
