package main

import (
	. "leetcode/utils"
)

// Given binary tree [3,9,20,null,null,15,7],
//	3
// / \
// 9  20
//	 /  \
//	15  7
// return its depth = 3
// 递归：遍历根节点的左孩子的高度和根节点右孩子的高度，取出两者的最大值 + 1
func maxDepth_dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 在递归过程中传递深度信息
	return Max(maxDepth_dfs(root.Left), maxDepth_dfs(root.Right)) + 1
}

// 层级遍历 借助队列实现
func maxDepth_bfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := []*TreeNode{root}
	level := 0 // 记录深度
	for len(q) > 0 {
		lq := len(q) // 当前层结点数
		for lq != 0 {
			node := q[0]
			q := q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			lq--
		}
		level++
	}
	return level
}
