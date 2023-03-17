package main

import . "leetcode/utils"

// 青蛙一次可以跳上1 或 2级台阶 求跳上 n 级台阶总共有多少种跳法
// f(n) 表示 n 级台阶的跳法 那上一步可能在第 n-1 或 n-2 级 分别需要再爬 1 或 2 级
// 得到 f(n) = f(n-1) + f(n-2) 可以由此式暴力 dfs
func dfs(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return dfs(n-1) + dfs(n-2)
}

func dfs_memo(n int) int {
	memo := make([]int, n+1)
	var dfs func(int) int
	dfs = func(n int) int {
		if n <= 2 {
			return n
		}
		if memo[n] != 0 {
			return memo[n]
		}
		memo[n] = dfs(n-1) + dfs(n-2)
		return memo[n]
	}
	return dfs(n)
}

// 动态规划
func numWays_dp(n int) int {
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}
	return dp[n]
}

func numWays_dp1(n int) int {
	pre, cur := 1, 1
	for i := 2; i <= n; i++ {
		pre, cur = cur%Mod, (pre+cur)%Mod
	}
	return cur
}
