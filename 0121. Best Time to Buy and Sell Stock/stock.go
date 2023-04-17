package main

import . "leetcode/utils"

// Input: prices = [7,1,5,3,6,4] Output: 5
// Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5
// 你只能选择某一天买入这只股票并选择在未来的某一个不同的日子卖出该股票
// 两个状态: 哪一天 & 是否持有股票 dp[i][0 or 1] 表示第 i 天是否持有股票(0 or 1)最终得到的利润
func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		// 不持有，可能有两个选择: hold or sell 对两者取最优
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i])
		// 持有，可能是 hold or buy
		dp[i][1] = Max(dp[i-1][1], -prices[i])
	}
	return dp[n-1][0]
}

// 单调栈 假如栈空或者入栈元素 > 栈顶元素直接入栈
// 假如入栈元素 < 栈顶元素则循环弹栈直到入栈元素 > 栈顶元素或栈空
// 在每次弹出的时候 拿它与买入的值 (也就是栈底) 做差 维护一个最大值
func maxProfit_stack(prices []int) int {
	ans, stack := 0, []int{}
	prices = append(prices, -1) // 新增哨兵
	for _, price := range prices {
		for len(stack) > 0 && stack[len(stack)-1] > price { // 维护单调栈
			ans = Max(ans, stack[len(stack)-1]-stack[0]) // 维护最大值
			stack = stack[:len(stack)-1] // 弹出
		}
		stack = append(stack, price)
	}
	return ans
}

func main() {
	maxProfit([]int{7, 1, 5, 3, 6, 4})
}
