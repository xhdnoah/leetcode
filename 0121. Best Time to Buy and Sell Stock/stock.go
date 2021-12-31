package main

import "fmt"

// Input: prices = [7,1,5,3,6,4] Output: 5
// Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
// Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
// 两个状态: 哪一天 & 是否持有股票 dp[i][0 or 1] 表示第 i 天是否持有股票(0 or 1)最终得到的利润
func maxProfit(prices []int) int {
	ln := len(prices)
	dp := make([][]int, ln)
	for i, price := range prices {
		dp[i] = make([]int, 2)
		if i-1 == -1 { // base case: 第 0 天
			dp[i][0] = 0
			dp[i][1] = -price
			continue
		}
		// 不持有，可能有两个选择: hold or sell 对两者取最优
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+price)
		// 持有，可能是 hold or buy
		dp[i][1] = max(dp[i-1][1], -price)
	}
	return dp[ln-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	maxProfit([]int{7, 1, 5, 3, 6, 4})
}
