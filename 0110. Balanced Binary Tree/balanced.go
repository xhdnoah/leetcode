package main

import . "leetcode/utils"

// 自底至顶：提前阻断
func isBalanced(root *TreeNode) bool {
	return dfs(root) != -1
}

// 深度优先中的后序遍历
func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left)
	if left == -1 {
		return -1
	}
	right := dfs(root.Right)
	if right == -1 {
		return -1
	}
	if Abs(left-right) < 2 {
		return Max(left, right) + 1
	} else {
		return -1
	}
}
