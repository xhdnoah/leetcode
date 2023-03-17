package main

import . "leetcode/utils"

type matrix [2][2]int

// 斐波那契数列 f(n+1) = f(n) + f(n-1)
// 递归：f(n) 问题拆分成 f(n-1) f(n-2) 两个子问题 会有重复的计算
// 尾递归优化：存储递归过程中 f(0) 到 f(n) 值 需要 O(N) 空间
// 动态规划：以 f(n+1) = f(n) + f(n-1) 为转移方程 时间复杂度 O(N)
func fib_dp(n int) int {
	if n < 2 {
		return n
	}
	p, q, r := 0, 0, 1 // 滚动数组
	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = (p + q) % Mod
	}
	return r
}

// 矩阵快速幂 O(logn)
func fib_pow(n int) int {
	if n < 2 {
		return n
	}
	res := pow(matrix{{1, 1}, {1, 0}}, n-1)
	return res[0][0]
}

func multiply(a, b matrix) (c matrix) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = (a[i][0]*b[0][j] + a[i][1]*b[1][j]) % Mod
		}
	}
	return
}

func pow(a matrix, n int) matrix {
	ret := matrix{{1, 0}, {0, 1}}
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			ret = multiply(ret, a)
		}
		a = multiply(a, a)
	}
	return ret
}
