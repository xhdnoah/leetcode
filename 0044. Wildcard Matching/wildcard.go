package main

// 参考 10.Regex Matching
func isMatch(s, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 1; i <= n; i++ {
		if p[i-1] != '*' { // 出现非 * 则退出循环
			break
		}
		dp[0][i] = true // s 连续为空则一定可以匹配
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			switch x := p[j-1]; {
			case x == '*':
				dp[i][j] = dp[i][j-1] || dp[i-1][j] // 使用或不使用 *
			case x == '?' || x == s[i-1]:
				dp[i][j] = dp[i-1][j-1] // 分别去掉当前的字节 从上一轮转移而来
			}
		}
	}
	return dp[m][n]
}
