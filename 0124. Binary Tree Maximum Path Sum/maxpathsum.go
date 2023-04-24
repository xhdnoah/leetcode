package main

import . "leetcode/utils"

// 后序遍历 难点在于返回什么
// dfs 递归应返回以根结点为端点的最大路径和而不是包含根结点的最大路径和
// 递归的返回值不是最终结果，需要使用一个全局变量在递归过程中动态更新其最大值
func maxPathSum(root *TreeNode) int {
	ans := -1 << 31 // 初始化为极小值
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		fromLeft := dfs(root.Left)   // 以左子树为端点的最大路径和
		fromRight := dfs(root.Right) // 以右子树为端点的最大路径和
		fromRoot := Maximum(
			root.Val,
			root.Val+fromLeft,
			root.Val+fromRight,
		)
		bypassRoot := root.Val + fromLeft + fromRight // 以 root 为中间节点的最大路径和
		ans = Maximum(ans, bypassRoot, fromRoot)
		return fromRoot // 返回以 root 为端点的最大路径和
	}
	dfs(root)
	return ans
}
