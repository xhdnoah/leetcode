package main

// Input: s = "babad" Output: "bab" 最长回文子串
// 动态规划 DP 时间复杂度 O(n^2)，空间复杂度 O(n^2)
// 定义 dp[i][j] 表示从字符串第 i 个字符到第 j 个字符这一段子串是否是回文串
// 由回文串的性质可得，回文串去掉一头一尾相同字符后，剩下的还是回文串
// 故状态转移方程 dp[i][j] = (s[i] == s[j]) && ((j-i < 3) || dp[i+1][j-1])
func longestPalindrome_dp(s string) string {
	res, dp := "", make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			dp[i][j] = (s[i] == s[j]) && ((j-i < 3) || dp[i+1][j-1])
			if dp[i][j] && (res == "" || j-i+1 > len(res)) {
				res = s[i : j+1]
			}
		}
	}
	return res
}
