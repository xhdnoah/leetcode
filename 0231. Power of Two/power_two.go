package main

// 一个数 n 是 2 的幂，当且仅当 n 是正整数，并且 n 的二进制表示中仅包含 1 个 1
// 使用位运算，提取 n 的二进制表示中最低位的那个 1 再判断剩余的数值是否为 0 即可
// n & (n-1) 可以直接将 n 二进制表示的最低位 1 移除
func isPowerOfTwo(num int) bool {
	return (num > 0 && ((num & (num - 1)) == 0))
}

// 数论: 在题目给定的 32 位有符号整数的范围内，最大的 2 的幂为
// 2^30 = 10737418242 只需要判断 n 是否是 2^30 的约数即可
func isPowerOfTwo1(num int) bool {
	return num > 0 && (1073741824%num == 0)
}
