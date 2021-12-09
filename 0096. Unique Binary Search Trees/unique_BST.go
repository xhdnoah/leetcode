package main

// Input: 3 Output: 5
// Explanation: Given n = 3, there are a total of 5 unique BST's:
//    1         3     3      2      1
//     \       /     /      / \      \
//      3     2     1      1   3      2
//     /     /       \                 \
//    2     1         2                 3
// 给定一个整数 n，求以 1 … n 为节点组成的二叉搜索树有多少种？
// DP: dp[n] 代表 1-n 个数能组成多少不同的二叉排序树
// F(i,n) 代表以 i 为根节点 1-n 个数组成不同的二叉排序树个数
// dp[n] = F(1,n)+F(2,n)+F(3,n)+...+F(n,n) dp[0]=dp[1]=1 而 F(i,n) = dp[i-1]*dp[n-i]
// 举例 [1,2,3,4,…,i,…,n-1,n]，以 i 为 根节点，那么左半边 [1,2,3,……,i-1] 和 右半边 [i+1,i+2,……,n-1,n]
// 分别能组成二叉排序树的不同个数相乘，即为以 i 为根节点，1-n 个数组成的二叉排序树的不同的个数，也即 F(i,n)
// 二叉排序树本身性质，右子树一定比左子树的值都大。故只需要根节点把树分成左右，不用关心左右两边数字大小，只需关心数字个数
// 状态转移方程 dp[n] = dp[0]*dp[n-1] + dp[1]*dp[n-2] + ... + dp[n-1]*dp[0]
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
