package main

import . "leetcode/utils"

// Input: [[1,3,1], [1,5,1], [4,2,1]] Output: 12 comes from 1->3->5->2->1
// DP: dp[i][j] 代表从棋盘左上角开始 到达单元格 (i,j) 时能拿到礼物的累计最大值
// 转移方程：i=0&j=0->起始元素；i=0&j!=0->第一行元素只可从左边到达；i!=0&j=0第一列元素只可从上边到达；i!=0&j!=0可从左或上边到达
func maxValue(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	dp := make([][]int, row)
	for x := range dp {
		dp[x] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			cur := grid[i][j]
			if i == 0 && j == 0 {
				dp[i][j] = cur
				continue
			}
			var tmp int
			if i == 0 {
				tmp = dp[i][j-1]
			} else if j == 0 {
				tmp = dp[j-1][j]
			} else {
				tmp = Max(dp[i][j-1], dp[i-1][j])
			}
			dp[i][j] = tmp + cur
		}
	}
	return dp[row-1][col-1]
}
