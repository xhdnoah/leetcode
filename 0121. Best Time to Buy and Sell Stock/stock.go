package main

import . "leetcode/utils"

// Input: prices = [7,1,5,3,6,4] Output: 5
// Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
// 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票
// 两个状态: 哪一天 & 是否持有股票 dp[i][0 or 1] 表示第 i 天是否持有股票(0 or 1)最终得到的利润
func maxProfit(prices []int) int {
	ln := len(prices)
	dp := make([][2]int, ln)
	dp[0][1] = -prices[0]
	for i := 1; i < ln; i++ {
		// 不持有，可能有两个选择: hold or sell 对两者取最优
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i])
		// 持有，可能是 hold or buy
		dp[i][1] = Max(dp[i-1][1], -prices[i])
	}
	return dp[ln-1][0]
}

func main() {
	maxProfit([]int{7, 1, 5, 3, 6, 4})
}
