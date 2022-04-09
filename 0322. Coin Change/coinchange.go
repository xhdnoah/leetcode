package main

import (
	"sort"
)

// Input: coins = [1,2,5], amount = 11 Output: 3 (11 = 5 + 5 + 1
// 动态规划 dp[i] = x 代表当目标金额为 i 至少需要 x 枚硬币
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0 // base case
	for i := 1; i < amount+1; i++ {
		dp[i] = amount + 1 // 初始化给一个极大值
	}
	for target := 1; target < len(dp); target++ {
		for _, coin := range coins {
			if coin <= target {
				// dp[target - coin] + 1(*coin)
				dp[target] = min(dp[target], 1+dp[target-coin])
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

// 遇到如 [1,1000000000,500000] 这样大值数组，DP 数组会浪费空间
// 用 DFS: 从左往右遍历降序 coins 遇到一些 amount 特别大的用例会很耗时
func coinChange_DFS(coins []int, amount int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(coins))) // 降序排序
	lc := len(coins)
	minus := coins[lc-1]
	ans := amount + 1
	var dfs func(int, int, int)
	cnt := 0
	dfs = func(start, amount, count int) { // DFS 搜索索引,剩余金额,已用硬币数量
		cnt++
		if amount == 0 {
			ans = min(ans, count)
			return
		}
		if start == lc {
			return
		}
		if minus > amount {
			return
		}

		coin := coins[start]
		division := amount / coin
		// 剪枝条件: 已用数量 + k 大于等于当前最小 ans
		for k := division; k >= 0 && count+k < ans; k-- {
			rem := amount - k*coin
			dfs(start+1, rem, count+k)
		}
	}
	dfs(0, amount, 0)
	if ans < amount+1 {
		return ans
	}
	return -1
}

//                       [5,2,1] 11, 0
//                  (k)2*5/ (k)1*5|
//                       /        |
//              [2,1] 1, 2    [2,1] 6,1
//                 0*2 /          |
//                    /       之后 amount 6 除以任何 coin 得到 k
//             [1] 1,2         所用硬币总数都大于第一个分支的 3 (此时的 ans)
//              1*1 /          故剪枝
//                 /
//           [] 0, 3

// 带备忘录 DFS 能使用 memo 的一个条件是其映射关系足够简单
func coinChange_Memo_DFS(coins []int, amount int) {
	infinity := amount + 1
	memo := make(map[int]int)
	var dfs func(int) int

	dfs = func(rem int) int {
		if rem < 0 {
			return infinity
		}
		if rem == 0 {
			return 0
		}
		if cnt, ok := memo[rem]; ok {
			return cnt
		}
		count := infinity
		for _, coin := range coins {
			if cnt := dfs(rem - coin); cnt != infinity {
				count = min(count, cnt+1)
			}
		}
		memo[rem] = count
		return count
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	coinChange_DFS([]int{411, 412, 413, 414, 415, 416, 417, 418, 419, 420, 421, 422}, 9864)
}
