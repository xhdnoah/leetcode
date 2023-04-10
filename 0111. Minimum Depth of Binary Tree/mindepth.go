package main

import (
	. "leetcode/utils"
)

// 最小深度是从根节点到最近叶子节点的最短路径上的节点数量
// BFS 利用队列记录广度优先扫描的当前层结点
func minDepth_BFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := []*TreeNode{root}
	depth := 1

	for len(q) > 0 {
		s := len(q)
		for i := 0; i < s; i++ {
			cur := q[0]
			q = q[1:]
			if cur.Left == nil && cur.Right == nil { // 叶子结点
				return depth
			}
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		depth++
	}
	return -1
}

// DFS 利用递归栈记录路径
// 对于每一个非叶子节点，只需要分别关注其左右子树的最小叶子节点深度
func minDepth(root *TreeNode) int {
	if root == nil { // base case
		return 0
	}
	left, right := minDepth(root.Left), minDepth(root.Right)
	// 关键在于理清楚递归结束条件
	// 如果左或右子树深度不为 0 即存在一个子树则当前子树最小深度就是该子树深度+1
	// 如果左右子树都不为 0 说明两颗子树都存在那么当前子树最小深度就是它们的较小值+1
	if left == 0 || right == 0 {
		return left + right + 1
	} else {
		return Min(left, right) + 1
	}
}
