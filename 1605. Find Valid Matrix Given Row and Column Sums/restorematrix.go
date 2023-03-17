package main

import . "leetcode/utils"

// 贪心算法
func restoreMatric(rowSum, colSum []int) [][]int {
	row, col := len(rowSum), len(colSum)
	matrix := make([][]int, row)
	for x := range matrix {
		matrix[x] = make([]int, col)
	}
	for i := range rowSum {
		for j := range colSum {
			m := Min(rowSum[i], colSum[j]) // 最大可填值
			matrix[i][j] = m
			rowSum[i] -= m // 更新和值
			colSum[j] -= m
		}
	}
	return matrix
}

func main() {
	var (
		rowSum = []int{5, 7, 10}
		colSum = []int{8, 6, 8}
	)
	restoreMatric(rowSum, colSum)
}
