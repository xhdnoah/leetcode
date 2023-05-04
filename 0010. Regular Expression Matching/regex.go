package main

// Given an input string 's' and a pattern 'p', implement regular expression matching with support for '.' and '*' where:
// '.' Matches any single character.​​​​
// '*' Matches zero or more of the preceding element.
// The matching should cover the entire input string (not partial).
// Input: s = "ab", p = ".*"
// Output: true
// Explanation: ".*" means "zero or more (*) of any character (.)".
 
// 关键在于理解特殊字符 '*' 的作用：匹配零个或多个前面的那一个元素，可以理解为前一个元素的消除或复制。
// f[i][j] 表示 s 的前 i 个字符与 p 中的前 j 个字符是否能够匹配。在进行状态转移时，我们考虑 p[j] 的匹配情况：
// 1. p[j] 是一个小写字母 那么必须在 s 中匹配一个相同小写字母 2. p[j] = '.'，则 p[j] 一定可以与 s[i] 匹配成功
// 3. p[j] = '*' 那么就表示可以对 p[j-1] 匹配任意自然数次, 以 s="abaaacd" p="aba*cd" s[i]='a' p[j−1]='a'和 p[j]='*' 为例
// 匹配 0 次的情况相当于在 p 中删去了 p[j−1] 和 p[j] 得转移方程 f[i][j] = f[i][j-2]
// 匹配 1 次意味着 p[j-1] p[j] 组成了 'a' 若 s[j]=p[j-1] 得转移方程 f[i][j] = f[i-1][j-2]
// 匹配 k 次意味着 p[j-1] p[j] 组成 k 个 'a' 若 s[i-k+1, ..., j]=p[j-1] 得转移方程 f[i][j] = f[i-k][j-2]
// 'a*' 这样的组合在匹配的过程中，本质只有两种：
// 一、匹配 s 末尾的 'a'，将 'a' 扔掉而 'a*' 还可继续进行匹配（递归匹配 >=1 次得 f[i][j] = s[i] = p[j-1] & f[i-1][j] 
// 二、不匹配字符，将 'a*' 扔掉不再进行匹配（匹配 0 次得 f[i][j] = f[i][j-2]
// 得到简化的转移方程 f[i][j] = f[i-1][j] & s[i] = p[j-1] or f[i][j-2], f[i][j] = f[i][j-2] & s[i] != p[j-1]
// 为减少对边界的判断 初始化 dp 数组维度为 (m+1) x (n+1) 具体代码实现时 s[i-1] p[j-1] 分别代表 s p 中第 i j 个字符
func isMatch(s, p string) bool {
	m, n := len(s), len(p)
	matches := func(i, j int) bool {
		return s[i-1] == p[j-1] || p[j-1] == '.'
	}
	dp := make([][]bool, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for j := 1; j <= n; j++ { // 不需要考虑 p 为空的情况 但要考虑 s 为空的情况
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}
	for i := 1; i <= m; i++ { 
		for j := 1; j <= n; j++ {
			if matches(i, j) {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				if matches(i, j-1) {
					dp[i][j] = dp[i][j-2] || dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-2]
				}
			}
		}
	}
	return dp[m][n]
}