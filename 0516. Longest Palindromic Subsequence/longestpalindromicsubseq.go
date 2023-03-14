package main

import . "leetcode/utils"

// Input: s = "bbbab" Output: 4
// One possible longest palindromic subsequence is "bbbb".
// DP: dp[i][j] 表示字符串 s 的下标范围 [i, j] 内的最长回文子序列的长度
func longestPalindromeSubseq(s string) int {
	ls := len(s)
	dp := make([][]int, ls)
	for i := range dp {
		dp[i] = make([]int, ls)
		dp[i][i] = 1 // base case
	}
	for i := ls - 1; i >= 0; i-- {
		for j := i + 1; j < ls; j++ {
			if s[i] == s[j] {
				// s[j]s[i]相等则先找到[i+1, j-1]范围内的最长回文串，再在首尾添加 i j 元素
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = Max(dp[i][j-1], dp[i+1][j])
			}
		}
	}
	return dp[0][ls-1]
}
