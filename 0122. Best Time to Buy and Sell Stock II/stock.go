package main

// 输入: prices = [1,2,3,4,5] 输出: 4
// 解释: 在第 1 天买入，在第 5 天卖出, 这笔交易所能获得利润 5-1 = 4
func maxProfit_dp(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

// 最大利润是所有上升区间之和
func maxProfit_greedy(prices []int) (ans int) {
	for i := 1; i < len(prices); i++ {
		ans += max(0, prices[i]-prices[i-1])
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
