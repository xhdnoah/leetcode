package utils

import "math"

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Minimum(nums ...int) int {
	m := math.MaxInt32
	for _, v := range nums {
		if v < m {
			m = v
		}
	}
	return m
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
