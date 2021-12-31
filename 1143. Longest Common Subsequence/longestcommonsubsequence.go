package main

// Input: text1 = "abcde", text2 = "ace" Output: 3
// The longest common subsequence is "ace" and its length is 3.
// dp[i+1][j+1] : 对于 s1[:i+1] 和 s2[:j+1] 它们的 LCS 长度
func longestCommonSubsequence(text1 string, text2 string) int {
	l1, l2 := len(text1), len(text2)
	dp := make([][]int, l1+1)
	for i := range dp {
		dp[i] = make([]int, l2+1)
	}
	for i, v1 := range text1 {
		for j, v2 := range text2 {
			if v1 == v2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[l1][l2]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
